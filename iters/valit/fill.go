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

// FillIterator represents an iterator yielding the same value over and over again.
type FillIterator[T any] struct{ v T }

func (it FillIterator[T]) Value() T   { return it.v }
func (it FillIterator[T]) Next() bool { return true }

// Fill provides an iterator which yields the same value over and over again.
func Fill[T any](v T) itkit.Iterator[T] {
	return &FillIterator[T]{v: v}
}
