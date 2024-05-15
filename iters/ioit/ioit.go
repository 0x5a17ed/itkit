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

package ioit

import (
	"sync/atomic"

	"github.com/0x5a17ed/itkit/iters/genit"
)

type GeneratorFn[T any] func(yield func(T)) error

type Generator[T any] struct {
	*genit.Generator[T]

	err atomic.Pointer[error]
}

// Close stops the generator and returns any error returned by
// the [GeneratorFn] function.
func (g *Generator[T]) Close() error {
	g.Generator.Stop()
	return g.Err()
}

// Err returns any error returned by the [GeneratorFn] function.
func (g *Generator[T]) Err() error {
	if err := g.err.Load(); err != nil {
		return *err
	}
	return nil
}

func Run[T any](fn GeneratorFn[T]) *Generator[T] {
	iog := &Generator[T]{}

	iog.Generator = genit.Run(func(yield func(T)) {
		err := fn(yield)
		iog.err.Store(&err)
	})

	return iog
}
