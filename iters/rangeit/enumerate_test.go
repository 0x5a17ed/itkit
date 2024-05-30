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

package rangeit_test

import (
	"testing"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/0x5a17ed/itkit/ittuple"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestEnumerate(t *testing.T) {
	fixture := func() itkit.Iterator[string] {
		return sliceit.In([]string{"A", "B", "C", "D"})
	}

	type testCase[I constraints.Integer, T any] struct {
		name string
		inp  itkit.Iterator[itlib.Pair[I, T]]
		want []itlib.Pair[I, T]
	}
	tt := []testCase[int, string]{
		{"empty", rangeit.Enumerate(itlib.Empty[string]()), nil},

		{"normal", rangeit.Enumerate(fixture()), []itlib.Pair[int, string]{
			ittuple.T2[int, string]{Left: 0, Right: "A"},
			ittuple.T2[int, string]{Left: 1, Right: "B"},
			ittuple.T2[int, string]{Left: 2, Right: "C"},
			ittuple.T2[int, string]{Left: 3, Right: "D"},
		}},

		{"from", rangeit.EnumerateFrom(1, fixture()), []itlib.Pair[int, string]{
			ittuple.T2[int, string]{Left: 1, Right: "A"},
			ittuple.T2[int, string]{Left: 2, Right: "B"},
			ittuple.T2[int, string]{Left: 3, Right: "C"},
			ittuple.T2[int, string]{Left: 4, Right: "D"},
		}},

		{"step", rangeit.EnumerateStep(0, 2, fixture()), []itlib.Pair[int, string]{
			ittuple.T2[int, string]{Left: 0, Right: "A"},
			ittuple.T2[int, string]{Left: 2, Right: "B"},
			ittuple.T2[int, string]{Left: 4, Right: "C"},
			ittuple.T2[int, string]{Left: 6, Right: "D"},
		}},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, sliceit.To(tc.inp))
		})
	}
}
