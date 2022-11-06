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

package rangeit

import (
	"golang.org/x/exp/constraints"

	"github.com/0x5a17ed/itkit"
)

// CountIterator represents an iterator which yields numbers starting
// at a given value and increases by a given value every time Next is
// called.
type CountIterator[T constraints.Integer] struct{ cur, next, step T }

func (c *CountIterator[T]) Value() T { return c.cur }

func (c *CountIterator[T]) Next() bool {
	c.cur, c.next = c.next, c.next+c.step
	return true
}

// Count returns an Iterator yielding numbers starting at 0 and
// increasing by 1.
func Count[T constraints.Integer]() itkit.Iterator[T] {
	return &CountIterator[T]{step: T(1)}
}

// CountFrom returns an Iterator yielding numbers starting at the
// given value in start and increasing by 1.
func CountFrom[T constraints.Integer](start T) itkit.Iterator[T] {
	return &CountIterator[T]{next: start, step: T(1)}
}

// CountStep returns an Iterator yielding numbers starting at the
// given value in start and increasing by the given value in step.
func CountStep[T constraints.Integer](start, step T) itkit.Iterator[T] {
	return &CountIterator[T]{next: start, step: step}
}
