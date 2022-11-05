// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io>
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

package funcit

import (
	"github.com/0x5a17ed/itkit"
)

// PullFnT represents a function that returns a single value and a
// boolean value indicating whenever calling the function will yield
// more values or not.
type PullFnT[T any] func() (T, bool)

// PullIterator represents an iterator which calls a function until
// it returns no more items.
type PullIterator[T any] struct {
	fn  PullFnT[T]
	cur T
}

func (it *PullIterator[T]) Next() (ok bool) { it.cur, ok = it.fn(); return }
func (it *PullIterator[T]) Value() T        { return it.cur }

// PullFn provides an iterator which calls the given function returning
// its items until the function signals there are no more items left.
func PullFn[T any](puller PullFnT[T]) itkit.Iterator[T] {
	return &PullIterator[T]{fn: puller}
}
