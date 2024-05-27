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

	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := itlib.Limit(10, itlib.Empty[any]())
		assert.False(t, it.Next())
	})

	t.Run("less", func(t *testing.T) {
		asserter := assert.New(t)

		src := rangeit.Range(10)

		it := itlib.Limit(4, src)

		asserter.Equal([]int{0, 1, 2, 3}, sliceit.To(it))

		asserter.False(it.Next())

		asserter.True(src.Next())
		asserter.Equal(4, src.Value())
	})

	t.Run("more", func(t *testing.T) {
		asserter := assert.New(t)

		src := rangeit.Range(4)

		it := itlib.Limit(10, src)

		asserter.Equal([]int{0, 1, 2, 3}, sliceit.To(it))

		asserter.False(it.Next())

		asserter.False(src.Next())
	})
}
