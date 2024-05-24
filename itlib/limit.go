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

// LimitIterator is an iterator yielding up to n items from the
// provided source iterator.
type LimitIterator[T any] struct {
	src itkit.Iterator[T]
	n   uint
}

// Ensure LimitIterator implements the iterator interface.
var _ itkit.Iterator[struct{}] = &LimitIterator[struct{}]{}

// Next implements the [itkit.Iterator.Next] interface.
func (it *LimitIterator[T]) Next() bool {
	if it.n == 0 || !it.src.Next() {
		return false
	}
	it.n -= 1
	return true
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *LimitIterator[T]) Value() (v T) {
	return it.src.Value()
}

func newLimitIterator[T any](n uint, src itkit.Iterator[T]) *LimitIterator[T] {
	return &LimitIterator[T]{n: n, src: src}
}

// Limit returns a new [LimitIterator] instance.
func Limit[T any](src itkit.Iterator[T], n uint) itkit.Iterator[T] {
	return newLimitIterator(n, src)
}
