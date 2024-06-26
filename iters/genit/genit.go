// Copyright (c) 2023 individual contributors.
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

// Package genit allows for stateful running goroutines to be used
// as iterators yielding values.
//
// Iterator functions:
//   - [Run] - yields items sent by a running goroutine
package genit

import (
	"github.com/0x5a17ed/coro"

	"github.com/0x5a17ed/itkit"
)

// YieldFn is a function that is called by a generator to send back
// values generated by the same generator.
type YieldFn[O any] func(O)

// GeneratorFn is a function that generates values which are sent back through
// the given [YieldFn] yield function.
type GeneratorFn[T any] func(yield func(T))

// Generator wraps a [GeneratorFn] in a [coro.C] coroutine and represents the
// receiving end.
type Generator[T any] struct {
	*coro.C[any, T]
	value T
}

// Next fetches the next value produced by the wrapped [GeneratorFn]
// and returns true whenever there is a new value available and false
// otherwise.
func (g *Generator[T]) Next() (ok bool) {
	g.value, ok = g.Resume(nil)
	return
}

// Value returns the latest value produced by the wrapped [GeneratorFn].
func (g *Generator[T]) Value() T {
	return g.value
}

// Iter returns the [Generator] as an [itkit.Iterator] value.
func (g *Generator[T]) Iter() itkit.Iterator[T] {
	return g
}

// Run starts the given [GeneratorFn] function as a new [Generator].
//
// The [Generator] will not be stopped automatically and [coro.C.Stop]
// must be called on the returned generator to stop it manually.
func Run[T any](fn GeneratorFn[T]) *Generator[T] {
	g := &Generator[T]{
		C: coro.NewSub[any, T](func(_ any, yield func(T) any) {
			fn(func(t T) { yield(t) })
		}),
	}

	return g
}
