// Copyright (c) 2022 individual contributors.
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

package valit

import (
	"github.com/0x5a17ed/itkit"
)

// Copier represents a struct that can copy itself.
type Copier[T any] interface {
	Copy() T
}

// CopyIterator represents an iterator yielding copies of a given
// value on every iteration of the iterator.
type CopyIterator[T Copier[T]] struct{ src, cur T }

func (it CopyIterator[T]) Value() T   { return it.cur }
func (it CopyIterator[T]) Next() bool { it.cur = it.src.Copy(); return true }

// Copies provides an iterator which yields copies of a given value on
// every iteration of the iterator.
func Copies[T Copier[T]](v T) itkit.Iterator[T] {
	return &CopyIterator[T]{src: v}
}
