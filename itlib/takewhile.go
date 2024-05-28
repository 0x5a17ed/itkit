// Copyright (c) 2024 individual contributors.
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

package itlib

import (
	"github.com/0x5a17ed/itkit"
)

type TakeWhileFn[T any] func(x T) bool

// TakeWhileIterator represents an iterator yielding items from a
// given source iterator until a given [TakeWhileFn] returns false.
type TakeWhileIterator[T any] struct {
	src itkit.Iterator[T]
	fn  TakeWhileFn[T]

	stopped bool
}

// Ensure TakeWhileIterator implements the iterator interface.
var _ itkit.Iterator[struct{}] = &TakeWhileIterator[struct{}]{}

// Next implements the [itkit.Iterator.Next] interface.
func (it *TakeWhileIterator[T]) Next() bool {
	if it.stopped {
		return false
	}

	if it.stopped = !it.src.Next(); it.stopped {
		return false
	}

	if it.stopped = !it.fn(it.src.Value()); it.stopped {
		return false
	}

	return true
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *TakeWhileIterator[T]) Value() (v T) {
	return it.src.Value()
}

// TakeWhile returns a new [TakeWhileIterator] value.
func TakeWhile[T any](src itkit.Iterator[T], fn TakeWhileFn[T]) itkit.Iterator[T] {
	return &TakeWhileIterator[T]{src: src, fn: fn}
}
