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

package itkit_test

import (
	"sort"
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit"
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
			got := itkit.ToSlice(itkit.Keys(tt.given).Iter())
			sort.Strings(got)
			assertpkg.Equalf(t, tt.wanted, got, "Keys(%v)", tt.given)
		})
	}
}
