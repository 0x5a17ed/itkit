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

// ChunkIterator yields iterators yielding up to n items from a source
// iterator until the source iterator is exhausted.
type ChunkIterator[T any] struct {
	src *PeekIterator[T]

	cur *LimitIterator[T]
	n   uint
}

// Ensure ChunkIterator implements the iterator interface.
var _ itkit.Iterator[itkit.Iterator[struct{}]] = &ChunkIterator[struct{}]{}

// Next implements the [itkit.Iterator.Next] interface.
func (it *ChunkIterator[T]) Next() bool {
	if it.cur != nil && it.cur.n > 0 {
		// Drop any items not consumed in the previous chunk
		// to make sure that the next chunk only contains
		// new items.
		Drop(it.cur.n, it.src.Iter())
	}

	if _, ok := it.src.Peek(); !ok {
		return false
	}
	it.cur = newLimitIterator(it.n, it.src.Iter())
	return true
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *ChunkIterator[T]) Value() itkit.Iterator[T] {
	return it.cur
}

// Chunk returns a new [ChunkIterator] value.
func Chunk[T any](n uint, src itkit.Iterator[T]) itkit.Iterator[itkit.Iterator[T]] {
	return &ChunkIterator[T]{src: newPeekIterator(src), n: n}
}
