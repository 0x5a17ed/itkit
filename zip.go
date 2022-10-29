// Copyright (c) 2022 individual contributors
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

type ZipIterator[T1, T2 any] struct {
	Left    Iterator[T1]
	Right   Iterator[T2]
	current Tuple2[T1, T2]
}

// Ensure ZipIterator conforms to the Iterator protocol.
var _ Iterator[Tuple2[struct{}, struct{}]] = &ZipIterator[struct{}, struct{}]{}

func (it *ZipIterator[T1, T2]) Next() bool {
	if !it.Left.Next() || !it.Right.Next() {
		it.current = Tuple2[T1, T2]{}
		return false
	}

	it.current = Tuple2[T1, T2]{it.Left.Value(), it.Right.Value()}
	return true
}

func (it *ZipIterator[T1, T2]) Value() Tuple2[T1, T2] {
	return it.current
}

// Zip returns an iterator that aggregates elements from the given iterators.
//
// The returned iterator yield Tuple2 values, where the i-th tuple contains
// the i-th element from each of the argument sequences or iterables.
// The iterator will stop when the shortest input iterable is exhausted.
func Zip[T1, T2 any](it1 Iterator[T1], it2 Iterator[T2]) Iterator[Tuple2[T1, T2]] {
	return &ZipIterator[T1, T2]{Left: it1, Right: it2}
}
