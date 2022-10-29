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

package runeit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/runeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
)

func TestString(t *testing.T) {
	t.Run("unicode", func(t *testing.T) {
		s := sliceit.To(runeit.InString("日本\x80語"))
		assert.Equal(t, []rune{0x65E5, 0x672C, 0xFFFD, 0x8A9E}, s)
	})

	t.Run("ascii", func(t *testing.T) {
		s := sliceit.To(runeit.InString("Hello World"))
		assert.Equal(t, []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0x6f, 0x72, 0x6c, 0x64}, s)
	})

	t.Run("late unicode", func(t *testing.T) {
		s := sliceit.To(runeit.InString("Hello Wörld"))
		assert.Equal(t, []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x57, 0xf6, 0x72, 0x6c, 0x64}, s)
	})
}
