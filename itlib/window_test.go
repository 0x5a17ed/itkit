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

package itlib_test

import (
	"testing"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/stretchr/testify/assert"
)

func TestWindow(t *testing.T) {
	p := func(n int) *int { return &n }

	fixture := func(x, y int) itkit.Iterator[*int] {
		return itlib.Map(rangeit.RangeFrom(x, y), func(el int) *int {
			return &el
		})
	}

	type testCase[T any] struct {
		name string
		it   itkit.Iterator[itkit.Iterator[T]]
		want [][]*int
	}
	tt := []testCase[*int]{
		{
			name: "empty",
			it:   itlib.Window(3, itlib.Empty[*int]()),
			want: nil,
		},
		{
			name: "golden",
			it:   itlib.Window(3, fixture(1, 6)),
			want: [][]*int{
				{p(1), p(2), p(3)},
				{p(2), p(3), p(4)},
				{p(3), p(4), p(5)},
			},
		},
		{
			name: "multistep",
			it: (&itlib.WindowIterator[*int]{
				Size:      3,
				Steps:     2,
				FillValue: nil,
				Source:    fixture(1, 6),
			}).Iter(),
			want: [][]*int{
				{p(1), p(2), p(3)},
				{p(3), p(4), p(5)},
			},
		},
		{
			name: "fill-start",
			it: (&itlib.WindowIterator[*int]{
				Size:      3,
				Steps:     1,
				FillValue: p(0),
				Source:    fixture(1, 3),
			}).Iter(),
			want: [][]*int{
				{p(1), p(2), p(0)},
			},
		},
		{
			name: "fill-end",
			it: (&itlib.WindowIterator[*int]{
				Size:      3,
				Steps:     2,
				FillValue: p(0),
				Source:    fixture(1, 7),
			}).Iter(),
			want: [][]*int{
				{p(1), p(2), p(3)},
				{p(3), p(4), p(5)},
				{p(5), p(6), p(0)},
			},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := sliceit.To(itlib.Map(tc.it, sliceit.To[*int]))

			assert.Exactly(t, tc.want, got)
		})
	}
}
