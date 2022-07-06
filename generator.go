// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package itkit

import (
	"errors"
	"runtime"

	"github.com/0x5a17ed/itkit/internal/event"
)

// ErrStopGenerator signals to the generator that it must stop.
var ErrStopGenerator = errors.New("stop generator")

// GeneratorFn is a function that generates items and sends them
// to the Iterator through G.Send.
type GeneratorFn[T any] func(g *G[T])

// G represents a Generator producing items which can be received
// through an iterator.
type G[T any] struct {
	ch      chan T
	closing event.E
	stopped event.E
	panic   any
}

// stop the generator function.
func (g *G[T]) stop() {
	// Synchronization point.  Tell the goroutine to stop and
	// wait for it until it has stopped.
	g.closing.Set()
	<-g.stopped.Wait()
}

// run the generator function.
func (g *G[T]) run(fn GeneratorFn[T]) {
	defer func() {
		if r := recover(); r != nil {
			rerr, ok := r.(error)
			if !ok || !errors.Is(rerr, ErrStopGenerator) {
				g.panic = r
			}
		}
		close(g.ch)
		g.stopped.Set()
	}()

	fn(g)
}

// Send is called from the Generator and sends the given Value to the
// Iterator caller.
//
// It panics when the Generator has to stop.
func (g *G[T]) Send(v T) {
	select {
	case <-g.closing.Wait():
		panic(ErrStopGenerator)
	case g.ch <- v:
	}
}

func newG[T any]() *G[T] {
	return &G[T]{ch: make(chan T)}
}

type GIterator[T any] struct {
	*ChannelIterator[T]
	g *G[T]
}

func (gs *GIterator[T]) Next() bool {
	if !gs.ChannelIterator.Next() {
		if gs.g.panic != nil {
			panic(gs.g.panic)
		}
		return false
	}
	return true
}

// Iter returns the underlying iterator of the generator yielding
// items produced by the generator until the generator function
// is gone.
func (gs *GIterator[T]) Iter() Iterator[T] { return gs }

// Stop stops the Generator. The Generator goroutine will be gone
// after this function completes.
func (gs *GIterator[T]) Stop() { gs.g.stop() }

// Generator returns the GeneratorFn as a new generator.
//
// The Generator API is experimental and probably will change.
func Generator[T any](fn GeneratorFn[T]) *GIterator[T] {
	gs := &GIterator[T]{g: newG[T]()}
	gs.ChannelIterator = Channel(gs.g.ch).(*ChannelIterator[T])
	runtime.SetFinalizer(gs, (*GIterator[T]).Stop)

	go gs.g.run(fn)

	return gs
}
