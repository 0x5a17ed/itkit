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

package valit_test

import (
	"testing"

	assertPkg "github.com/stretchr/testify/assert"
	requirePkg "github.com/stretchr/testify/require"

	"github.com/0x5a17ed/itkit/iters/valit"
)

func TestFillIterator(t *testing.T) {
	type object struct{}

	t.Run("integer", func(t *testing.T) {
		iter := valit.Fill(10)

		var got []int
		for i := 0; i < 5; i += 1 {
			requirePkg.True(t, iter.Next())
			got = append(got, iter.Value())
		}
		assertPkg.Equal(t, []int{10, 10, 10, 10, 10}, got)
	})

	t.Run("struct", func(t *testing.T) {
		var o object
		iter := valit.Fill(&o)
		for i := 0; i < 5; i += 1 {
			requirePkg.True(t, iter.Next())
			assertPkg.Same(t, &o, iter.Value())
		}
	})
}
