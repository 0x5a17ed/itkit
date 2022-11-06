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

package sliceit_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestSlice(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		s := sliceit.To(rangeit.Range(10))
		assertpkg.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
	})

	t.Run("structs-values", func(t *testing.T) {
		type foo struct{ v int }

		s := sliceit.To(itlib.Map(rangeit.Range(5), func(v int) foo { return foo{v: v} }))
		assertpkg.Equal(t, []foo{{0}, {1}, {2}, {3}, {4}}, s)
	})
}

func TestSliceIterator(t *testing.T) {
	it := sliceit.In([]int{1, 2, 3})

	var values []int
	for it.Next() {
		values = append(values, it.Value())
	}

	assertpkg.Equal(t, []int{1, 2, 3}, values)
}
