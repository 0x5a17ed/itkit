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

type RangeIter[T constraints.Signed] struct {
	index, start, step, length, current T
}

func (r *RangeIter[T]) Next() bool {
	if r.index >= r.length {
		return false
	}
	r.current = r.start + r.step*r.index
	r.index += 1
	return true
}

func (r RangeIter[T]) Value() T { return r.current }

type RangeOptions struct{ start, step Option[int] }

type OptionFn func(opts *RangeOptions)

func WithStart(v int) OptionFn { return func(opts *RangeOptions) { opts.start.Set(v) } }
func WithStep(v int) OptionFn  { return func(opts *RangeOptions) { opts.step.Set(v) } }

// Range returns an iterator yielding [start, stop)
func Range[T constraints.Signed](stop T, optFuncs ...OptionFn) Iterator[T] {
	var opts RangeOptions
	for _, opt := range optFuncs {
		opt(&opts)
	}

	startVal, stepVal := T(opts.start.OrElse(0)), T(opts.step.OrElse(1))

	var lo, hi T
	if stepVal > 0 {
		lo, hi = startVal, stop
	} else {
		lo, hi, stepVal = stop, startVal, -1*stepVal
	}

	var length T
	if hi > lo {
		length = (((hi - lo) - 1) / stepVal) + 1
	}
	return &RangeIter[T]{start: startVal, step: T(opts.step.OrElse(1)), length: length}
}
