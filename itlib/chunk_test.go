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
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	tt := []struct {
		name string
		inp  []string
		tx   func(el itkit.Iterator[string]) itkit.Iterator[string]
		exp  [][]string
	}{
		{
			name: "emtpy",
			inp:  nil,
			exp:  nil,
		},

		{
			name: "partial-first",
			inp:  []string{"A", "B"},
			exp:  [][]string{{"A", "B"}},
		},

		{
			name: "complete",
			inp: []string{
				"A", "B", "C", "D", "E", "F",
			},
			exp: [][]string{
				{"A", "B", "C"},
				{"D", "E", "F"},
			},
		},

		{
			name: "partial-last",
			inp: []string{
				"A", "B", "C", "D", "E", "F", "G",
			},
			exp: [][]string{
				{"A", "B", "C"},
				{"D", "E", "F"},
				{"G"},
			},
		},

		{
			name: "half-consumed",
			inp: []string{
				"A", "B", "C", "D", "E", "F", "G",
			},
			tx: func(el itkit.Iterator[string]) itkit.Iterator[string] {
				return itlib.Limit(2, el)
			},
			exp: [][]string{
				{"A", "B"},
				{"D", "E"},
				{"G"},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			it := itlib.Chunk(3, sliceit.In(tc.inp))

			if tc.tx != nil {
				it = itlib.Map(it, tc.tx)
			}

			got := sliceit.To(itlib.Map(it, sliceit.To[string]))

			assert.Equal(t, tc.exp, got)
		})
	}
}
