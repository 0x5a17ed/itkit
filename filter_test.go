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
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit"
)

func TestFilter(t *testing.T) {
	t.Run("FilterRange", func(t *testing.T) {
		assert := assertpkg.New(t)
		s := itkit.Slice(itkit.Filter(itkit.Range(10), func(i int) bool { return i%2 == 0 }))
		assert.Equal([]int{0, 2, 4, 6, 8}, s)
	})

	t.Run("FilterString", func(t *testing.T) {
		assert := assertpkg.New(t)
		s := itkit.String(itkit.Filter(itkit.Runes("Hello Wörld"), func(c rune) bool { return 'a' <= c && c <= 'z' }))
		assert.Equal("ellorld", s)
	})
}
