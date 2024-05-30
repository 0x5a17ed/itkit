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

package rangeit

import (
	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/itlib"
	"golang.org/x/exp/constraints"
)

// EnumerateStep returns an iterator that yields pairs containing the
// index, of the item in the source iterator, starting at a given value
// which is incremented at a given step value and the item itself.
func EnumerateStep[T any, I constraints.Integer](
	start, step I,
	src itkit.Iterator[T],
) itkit.Iterator[itlib.Pair[I, T]] {
	return itlib.Zip(CountStep(start, step), src)
}

// EnumerateFrom returns an iterator that yields pairs containing the
// index, of the item in the source iterator, starting at a given value
// and the item itself.
func EnumerateFrom[T any, I constraints.Integer](start I, src itkit.Iterator[T]) itkit.Iterator[itlib.Pair[I, T]] {
	return itlib.Zip(CountFrom(start), src)
}

// Enumerate returns an iterator that yields pairs containing the index
// of the item in the source iterator and the item itself.
func Enumerate[T any](src itkit.Iterator[T]) itkit.Iterator[itlib.Pair[int, T]] {
	return itlib.Zip(Count[int](), src)
}
