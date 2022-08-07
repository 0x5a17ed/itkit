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

// Send is called from a running [GeneratorFn] and sends the given Value to the
// [Iterator] caller.
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

// GIterator represents the receiver facing [Iterator] end of a Generator.
type GIterator[T any] struct {
	*ChannelIterator[T]
	g *G[T]
}

// Next fetches the next Item produced by the Generator and returns true
// whenever there is a new item available and false otherwise.
func (gi *GIterator[T]) Next() bool {
	if !gi.ChannelIterator.Next() {
		if gi.g.panic != nil {
			panic(gi.g.panic)
		}
		return false
	}
	return true
}

// Iter returns the underlying [Iterator] of the [GIterator] yielding
// items produced by the generator until the generator function
// is gone.
func (gi *GIterator[T]) Iter() Iterator[T] { return gi }

// Stop stops the Generator.
//
// The Generator Function will be stopped through a panic call that is
// meant to kill the running generator.
//
// The Generator goroutine will be gone after this function completes.
func (gi *GIterator[T]) Stop() { gi.g.stop() }

// GeneratorNoGC starts the [GeneratorFn] function as a new [GIterator].
//
// Please consider using [Generator] for simplicity reasons unless
// the Garbage Collector is a concern.
//
// The Generator will not be stopped automatically and [itkit.GIterator.Stop]
// must be called on the returned  to stop the generator manually.
//
// The Generator API is experimental and probably will change.
func GeneratorNoGC[T any](fn GeneratorFn[T]) *GIterator[T] {
	gi := &GIterator[T]{g: newG[T]()}
	gi.ChannelIterator = Channel(gi.g.ch).(*ChannelIterator[T])
	go gi.g.run(fn)

	return gi
}

// Generator starts the [GeneratorFn] function as a new [GIterator].
//
// The Generator will be stopped automatically when the returned
// [GIterator] is garbage collected.
//
// The Generator API is experimental and probably will change.
func Generator[T any](fn GeneratorFn[T]) *GIterator[T] {
	gi := GeneratorNoGC(fn)
	runtime.SetFinalizer(gi, (*GIterator[T]).Stop)
	return gi
}
