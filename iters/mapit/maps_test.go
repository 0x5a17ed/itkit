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

package mapit_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"

	"github.com/0x5a17ed/itkit/iters/mapit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/0x5a17ed/itkit/ittuple"
)

func TestKeys_String(t *testing.T) {
	tests := []struct {
		name   string
		given  map[string]int
		wanted []string
	}{
		{"nil map", nil, nil},
		{"empty map", map[string]int{}, nil},

		{"key set ABC", map[string]int{"A": 1, "B": 2, "C": 3}, []string{"A", "B", "C"}},
		{"key set XYZ", map[string]int{"X": 1, "Y": 2, "Z": 3}, []string{"X", "Y", "Z"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := mapit.Keys(tt.given)
			defer g.Stop()

			got := sliceit.To(g.Iter())
			sort.Strings(got)
			assert.Equalf(t, tt.wanted, got, "Keys(%v)", tt.given)
		})
	}
}

func TestMapValues(t *testing.T) {
	tests := []struct {
		name   string
		given  map[string]int
		wanted []int
	}{
		{"nil map", nil, nil},
		{"empty map", map[string]int{}, nil},

		{"key set ABC", map[string]int{"A": 1, "B": 2, "C": 3}, []int{1, 2, 3}},
		{"key set ZYX", map[string]int{"Z": 11, "Y": 13, "X": 17}, []int{11, 13, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := mapit.Values(tt.given)
			defer g.Stop()

			got := sliceit.To(g.Iter())
			sort.Ints(got)
			assert.Equalf(t, tt.wanted, got, "Values(%v)", tt.given)
		})
	}
}

func TestInMap(t *testing.T) {
	g := mapit.In(map[string]int{
		"foo": 23,
		"baa": 42,
		"baz": 17,
	})
	defer g.Stop()

	s := sliceit.To(g.Iter())

	slices.SortFunc(s, func(a, b itlib.Pair[string, int]) int {
		al, _ := a.Values()
		bl, _ := b.Values()

		return strings.Compare(al, bl)
	})

	assert.Equal(t, []itlib.Pair[string, int]{
		ittuple.T2[string, int]{"baa", 42},
		ittuple.T2[string, int]{"baz", 17},
		ittuple.T2[string, int]{"foo", 23},
	}, s)
}

func TestToMap(t *testing.T) {
	m := mapit.To(sliceit.In([]itlib.Pair[string, int]{
		ittuple.T2[string, int]{"baa", 42},
		ittuple.T2[string, int]{"baz", 17},
		ittuple.T2[string, int]{"foo", 23},
	}))

	assert.Equal(t, map[string]int{
		"foo": 23,
		"baa": 42,
		"baz": 17,
	}, m)
}
