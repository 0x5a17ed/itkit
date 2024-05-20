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

// CycleIterator represents an iterator yielding items from a given
// source iterable and saving a copy of each returned item.  When the
// given source iterable is exhausted, returns items from the saved
// list. The iterable repeats the process forever.
type CycleIterator[T any] struct {
	src itkit.Iterator[T]

	cur       T
	repeating bool
	copies    []T
	index     int
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *CycleIterator[T]) Value() T {
	return it.cur
}

// Next implements the [itkit.Iterator.Next] interface.
func (it *CycleIterator[T]) Next() bool {
	if !it.repeating {
		if it.src.Next() {
			it.cur = it.src.Value()
			it.copies = append(it.copies, it.cur)
			return true
		}
		it.repeating = true
	}

	if len(it.copies) == 0 {
		return false
	}

	it.cur, it.index = it.copies[it.index], it.index+1
	if it.index >= len(it.copies) {
		it.index = 0
	}
	return true
}

// Cycle returns a new [CycleIterator] value.
func Cycle[T any](src itkit.Iterator[T]) itkit.Iterator[T] {
	return &CycleIterator[T]{src: src}
}
