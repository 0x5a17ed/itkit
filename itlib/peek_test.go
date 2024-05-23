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

func TestPeek(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		it := itlib.Peek(itlib.Empty[int]())

		_, ok := it.Peek()
		assert.False(t, ok)

		assert.False(t, it.Next())
	})

	t.Run("one", func(t *testing.T) {
		asserter := assert.New(t)

		it := itlib.Peek(rangeit.RangeFrom(1, 3))

		// Peeking into the iterator yields the first item.
		v, ok := it.Peek()
		asserter.True(ok)
		asserter.Equal(1, v)

		// Peeking again yields the same value as before.
		v, ok = it.Peek()
		asserter.True(ok)
		asserter.Equal(1, v)

		// Advancing the iterator yields the same value again.
		asserter.True(it.Next())
		asserter.Equal(1, v)

		// Peeking again yields the next value in the source iterator.
		v, ok = it.Peek()
		asserter.True(ok)
		asserter.Equal(2, v)
	})

	t.Run("mirror", func(t *testing.T) {
		asserter := assert.New(t)

		it := itlib.Peek(rangeit.RangeFrom(1, 5))

		// Iterating the Peek iterator yields all items normally.
		asserter.Equal([]int{1, 2, 3, 4}, sliceit.To(it.Iter()))
	})
}
