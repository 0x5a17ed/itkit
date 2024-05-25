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

// PeekIterator represents an iterator allowing to peek the next item
// before advancing to it.
type PeekIterator[T any] struct {
	src itkit.Iterator[T]

	cur    T
	cached T
	has    bool
}

// Ensure PeekIterator implements the iterator interface.
var _ itkit.Iterator[struct{}] = &PeekIterator[struct{}]{}

// Next implements the [itkit.Iterator.Next] interface.
func (it *PeekIterator[T]) Next() (ok bool) {
	if ok = it.has; ok {
		// Try to consume the cached item first.
		it.cur, it.has = it.cached, false
	} else if ok = it.src.Next(); ok {
		// Try to get an item from the source iterator.
		it.cur = it.src.Value()
	}
	return
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *PeekIterator[T]) Value() T {
	return it.cur
}

// Peek returns the next item without advancing the iterator.
//
// Advances the source iterator to the next item only if necessary.
func (it *PeekIterator[T]) Peek() (v T, ok bool) {
	if it.has {
		return it.cached, true
	}

	if it.has = it.src.Next(); it.has {
		it.cached = it.src.Value()
	}
	return it.cached, it.has
}

// Iter returns the [PeekIterator] as an [itkit.Iterator] value.
func (it *PeekIterator[T]) Iter() itkit.Iterator[T] {
	return it
}

func newPeekIterator[T any](src itkit.Iterator[T]) *PeekIterator[T] {
	return &PeekIterator[T]{src: src}
}

// Peek returns an Iterator that is always exhausted.
func Peek[T any](src itkit.Iterator[T]) *PeekIterator[T] {
	return newPeekIterator(src)
}
