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

func TestApply(t *testing.T) {
	assert := assertpkg.New(t)

	var loops, output int
	itkit.Apply(itkit.Range(10), func(item int) {
		loops += 1
		output += item
		return
	})

	assert.Equal(10, loops)
	assert.Equal(45, output)
}

func TestApplyN(t *testing.T) {
	assert := assertpkg.New(t)

	var runs, output int
	itkit.ApplyN(itkit.Range(10), func(index, item int) {
		runs += 1
		output += index + item
		return
	})

	assert.Equal(10, runs)
	assert.Equal(90, output)
}

func TestEach(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		assert := assertpkg.New(t)

		var loops, output int
		itkit.Each(itkit.Range(10), func(item int) (o bool) {
			loops += 1
			output += item
			return
		})

		assert.Equal(10, loops)
		assert.Equal(45, output)
	})

	t.Run("half", func(t *testing.T) {
		assert := assertpkg.New(t)

		var loops, output int
		itkit.Each(itkit.Range(10), func(item int) (o bool) {
			loops += 1
			output += item
			return item > 3
		})

		assert.Equal(5, loops)
		assert.Equal(10, output)
	})
}

func TestEachIndex(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		assert := assertpkg.New(t)

		var runs, output int
		itkit.EachN(itkit.Range(10), func(index, item int) (o bool) {
			runs += 1
			output += index + item
			return
		})

		assert.Equal(10, runs)
		assert.Equal(90, output)
	})

	t.Run("half", func(t *testing.T) {
		assert := assertpkg.New(t)

		var runs, output int
		itkit.EachN(itkit.Range(10), func(index, item int) (o bool) {
			runs += 1
			output += index + item
			return index >= 4
		})

		assert.Equal(5, runs)
		assert.Equal(20, output)
	})
}

func TestReduce(t *testing.T) {
	var n int

	n = itkit.ReduceWithInitial(5, itkit.InSlice([]int{1, 2, 3, 4, 5}), func(a, b int) int { return a * b })
	assertpkg.Equal(t, 600, n)

	n = itkit.Reduce(itkit.InSlice([]int{1, 2, 3, 4, 5}), func(a, b int) int { return a*2 + b })
	assertpkg.Equal(t, 57, n)
}

func TestSum(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		var n float64

		n = itkit.SumWithInitial(1, itkit.InSlice([]float64{1.1, 2.2, 3.3}))
		assertpkg.LessOrEqual(t, n-7.6, 0.1)

		n = itkit.Sum(itkit.InSlice([]float64{1.1, 2.2, 3.3}))
		assertpkg.LessOrEqual(t, n-6.6, 0.1)
	})

	t.Run("int", func(t *testing.T) {
		var n int

		n = itkit.SumWithInitial(1, itkit.InSlice([]int{-1, -2, -3}))
		assertpkg.Equal(t, -5, n)

		n = itkit.Sum(itkit.InSlice([]int{-1, -2, -3}))
		assertpkg.Equal(t, -6, n)
	})

	t.Run("uint", func(t *testing.T) {
		var n uint

		n = itkit.SumWithInitial(1, itkit.InSlice([]uint{1, 2, 3}))
		assertpkg.Equal(t, uint(7), n)

		n = itkit.Sum(itkit.InSlice([]uint{1, 2, 3}))
		assertpkg.Equal(t, uint(6), n)
	})

	t.Run("string", func(t *testing.T) {
		var n string

		n = itkit.SumWithInitial("def", itkit.InSlice([]string{"a", "b", "c"}))
		assertpkg.Equal(t, "defabc", n)

		n = itkit.Sum(itkit.InSlice([]string{"a", "b", "c"}))
		assertpkg.Equal(t, "abc", n)
	})
}
