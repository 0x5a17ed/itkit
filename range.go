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

package itkit

import (
	"golang.org/x/exp/constraints"
)

// RangeIter provides an [Iterator] over an immutable sequence of numbers.
type RangeIter[T constraints.Signed] struct {
	index, start, step, length, current T
}

// Next advances the iterator to the first/next item,
// returning true if successful meaning there is an item
// available to be fetched with Value and false otherwise.
//
// See [Iterator] for more details.
func (r *RangeIter[T]) Next() bool {
	if r.index >= r.length {
		return false
	}
	r.current = r.start + r.step*r.index
	r.index += 1
	return true
}

func (r *RangeIter[T]) Value() T { return r.current }

func newRange[T constraints.Signed](start, stop, step T) Iterator[T] {
	stepArg := step

	var lo, hi T
	if step > 0 {
		lo, hi = start, stop
	} else {
		lo, hi, step = stop, start, -1*step
	}

	var length T
	if hi > lo {
		length = (((hi - lo) - 1) / step) + 1
	}

	return &RangeIter[T]{start: start, step: stepArg, length: length}
}

// Range returns an iterator yielding [0 .. stop)
func Range[T constraints.Signed](stop T) Iterator[T] {
	return newRange(0, stop, 1)
}

// RangeFrom returns an iterator yielding [start .. stop)
func RangeFrom[T constraints.Signed](start, stop T) Iterator[T] {
	return newRange(start, stop, 1)
}

// RangeSteps returns an iterator yielding [start .. start+step*n .. stop)
func RangeSteps[T constraints.Signed](start, stop, step T) Iterator[T] {
	return newRange(start, stop, step)
}
