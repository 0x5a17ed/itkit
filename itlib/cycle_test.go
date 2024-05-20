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

	"github.com/0x5a17ed/itkit/iters/runeit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := itlib.Cycle(itlib.Empty[int]())

		assert.False(t, it.Next())
	})

	tt := []struct {
		name string
		inp  string
	}{
		{"single", "A"},
		{"multiple", "ABCD"},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			asserter := assert.New(t)

			inp := []rune(tc.inp)

			it := itlib.Cycle[rune](runeit.InString(tc.inp))
			for i := 0; i < len(inp)*2; i++ {
				v := itlib.HeadOrElse(it, rune(0))
				asserter.Equal(v, inp[i%len(tc.inp)])
			}
		})
	}
}
