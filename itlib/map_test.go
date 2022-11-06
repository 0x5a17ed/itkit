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

package itlib_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/runeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestMap(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		s := sliceit.To(itlib.Map(runeit.InString("Hello World!"), func(r rune) string { return string(r) }))
		assertpkg.Equal(t, []string{"H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d", "!"}, s)
	})

	t.Run("integers", func(t *testing.T) {
		s := sliceit.To(itlib.Map(rangeit.Range(4), func(v int) int { return 1 << v }))
		assertpkg.Equal(t, []int{1, 2, 4, 8}, s)
	})
}
