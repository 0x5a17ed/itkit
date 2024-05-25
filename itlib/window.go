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

type windowSubIterator[T any] struct{ parent *WindowIterator[T] }

func (it windowSubIterator[T]) Next() bool { return it.parent.windowNext() }
func (it windowSubIterator[T]) Value() T   { return it.parent.windowValue() }

// WindowIterator
type WindowIterator[T any] struct {
	// Size specifies the window size.
	Size uint

	// Steps specifies the number of items the iterator will
	// advance for each call of the Next function.
	Steps uint

	// FillValue is used to supplement missing items in case the
	// window is larger than the iterable can yield items.
	FillValue T

	// Source is the original source to yield items from.
	Source itkit.Iterator[T]

	window []T
	offset uint
	index  uint
	cur    T
}

// Ensure WindowIterator implements the iterator interface.
var _ itkit.Iterator[itkit.Iterator[struct{}]] = &WindowIterator[struct{}]{}

func (it *WindowIterator[T]) windowNext() (ok bool) {
	if ok = it.index < it.Size; ok {
		it.cur, it.index = it.window[(it.offset+it.index)%it.Size], it.index+1
	}
	return
}

func (it *WindowIterator[T]) windowValue() T {
	return it.cur
}

func (it *WindowIterator[T]) fillWindow(n uint) bool {
	if !it.Source.Next() {
		return false
	}

	it.window[it.offset] = it.Source.Value()
	for i := uint(1); i < n; i++ {
		j := (it.offset + i) % it.Size
		if it.Source.Next() {
			it.window[j] = it.Source.Value()
		} else {
			it.window[j] = it.FillValue
		}
	}

	return true
}

// Next implements the [itkit.Iterator.Next] interface.
func (it *WindowIterator[T]) Next() (ok bool) {
	if it.window == nil {
		it.window = make([]T, it.Size)
		if ok = it.fillWindow(it.Size); !ok {
			it.window = nil
		}
		return
	}

	if ok = it.fillWindow(it.Steps); ok {
		it.offset, it.index = (it.offset+it.Steps)%it.Size, 0
	}

	return
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *WindowIterator[T]) Value() itkit.Iterator[T] {
	return windowSubIterator[T]{parent: it}
}

// Iter returns the [WindowIterator] as an [itkit.Iterator] value.
func (it *WindowIterator[T]) Iter() itkit.Iterator[itkit.Iterator[T]] {
	return it
}

// Window returns a new [WindowIterator] value.
func Window[T any](n uint, src itkit.Iterator[T]) itkit.Iterator[itkit.Iterator[T]] {
	var zero T
	return &WindowIterator[T]{
		Size:      n,
		Steps:     1,
		FillValue: zero,
		Source:    src,
	}
}
