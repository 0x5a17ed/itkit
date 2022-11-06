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

package funcit

import (
	"github.com/0x5a17ed/itkit"
)

// NextFnT advances the iterator to the first/next item, returning true
// if successful meaning there is an item  available to be fetched and
// false otherwise.
type NextFnT func() bool

// ValueFnT returns the current item of the iterator.
type ValueFnT[T any] func() T

// StateIterator represents an iterator which calls a given next
// function to test whenever there is another item available and
// yields the values returned from a provided value retrieval
// function.
type StateIterator[T any] struct {
	n NextFnT
	v ValueFnT[T]
}

func (it StateIterator[T]) Value() T   { return it.v() }
func (it StateIterator[T]) Next() bool { return it.n() }

// IterFn provides an iterator which calls the provided next function
// to test whenever there is another item available and yields the
// values returned from the provided value function.
func IterFn[T any](nextFn NextFnT, valueFn ValueFnT[T]) itkit.Iterator[T] {
	return &StateIterator[T]{n: nextFn, v: valueFn}
}
