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

	"github.com/0x5a17ed/itkit"
	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestApply(t *testing.T) {
	assert := assertpkg.New(t)

	var loops, output int
	itlib.Apply(rangeit.Range(10), func(item int) {
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
	itlib.ApplyN(rangeit.Range(10), func(index, item int) {
		runs += 1
		output += index + item
		return
	})

	assert.Equal(10, runs)
	assert.Equal(90, output)
}

func TestApplyTo(t *testing.T) {
	assert := assertpkg.New(t)

	var s []int
	itlib.ApplyTo(rangeit.Range(10), &s, func(s *[]int, v int) { *s = append(*s, v) })

	assert.Equal([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s)
}

func TestEach(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		assert := assertpkg.New(t)

		var loops, output int
		itlib.Each(rangeit.Range(10), func(item int) (o bool) {
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
		itlib.Each(rangeit.Range(10), func(item int) (o bool) {
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
		itlib.EachN(rangeit.Range(10), func(index, item int) (o bool) {
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
		itlib.EachN(rangeit.Range(10), func(index, item int) (o bool) {
			runs += 1
			output += index + item
			return index >= 4
		})

		assert.Equal(5, runs)
		assert.Equal(20, output)
	})
}

func TestReduce(t *testing.T) {
	n := itlib.Reduce(sliceit.In([]int{1, 2, 3, 4, 5}), func(a, b int) int { return a*2 + b })
	assertpkg.Equal(t, 57, n)
}

func TestFold(t *testing.T) {
	n := itlib.Fold(5, sliceit.In([]int{1, 2, 3, 4, 5}), func(a, b int) int { return a * b })
	assertpkg.Equal(t, 600, n)
}

func TestSum(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		var n float64

		n = itlib.SumWithInitial(1, sliceit.In([]float64{1.1, 2.2, 3.3}))
		assertpkg.LessOrEqual(t, n-7.6, 0.1)

		n = itlib.Sum(sliceit.In([]float64{1.1, 2.2, 3.3}))
		assertpkg.LessOrEqual(t, n-6.6, 0.1)
	})

	t.Run("int", func(t *testing.T) {
		var n int

		n = itlib.SumWithInitial(1, sliceit.In([]int{-1, -2, -3}))
		assertpkg.Equal(t, -5, n)

		n = itlib.Sum(sliceit.In([]int{-1, -2, -3}))
		assertpkg.Equal(t, -6, n)
	})

	t.Run("uint", func(t *testing.T) {
		var n uint

		n = itlib.SumWithInitial(1, sliceit.In([]uint{1, 2, 3}))
		assertpkg.Equal(t, uint(7), n)

		n = itlib.Sum(sliceit.In([]uint{1, 2, 3}))
		assertpkg.Equal(t, uint(6), n)
	})

	t.Run("string", func(t *testing.T) {
		var n string

		n = itlib.SumWithInitial("def", sliceit.In([]string{"a", "b", "c"}))
		assertpkg.Equal(t, "defabc", n)

		n = itlib.Sum(sliceit.In([]string{"a", "b", "c"}))
		assertpkg.Equal(t, "abc", n)
	})
}

func TestDrop(t *testing.T) {
	type args[T any] struct {
		n  uint
		it itkit.Iterator[T]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tt := []testCase[int]{
		{"empty", args[int]{4, itlib.Empty[int]()}, nil},
		{"none", args[int]{0, rangeit.Range(5)}, []int{0, 1, 2, 3, 4}},
		{"all", args[int]{5, rangeit.Range(5)}, nil},
		{"less", args[int]{3, rangeit.Range(5)}, []int{3, 4}},
	}
	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceit.To(itlib.Drop(tt.args.n, tt.args.it))
			assertpkg.Equalf(t, tt.want, got, "Drop(%v, ...)", tt.args.n)
		})
	}
}

func TestAny(t *testing.T) {
	type testCase[T any] struct {
		name   string
		args   []bool
		wantOk bool
	}
	tt := []testCase[bool]{
		{"false", []bool{false, false, false}, false},
		{"true", []bool{false, false, true}, true},
	}
	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			got := itlib.Any(sliceit.In(tt.args), func(item bool) bool {
				return item
			})
			assertpkg.Equal(t, tt.wantOk, got)
		})
	}
}

func TestAll(t *testing.T) {
	type testCase[T any] struct {
		name   string
		args   []bool
		wantOk bool
	}
	tt := []testCase[bool]{
		{"false", []bool{false, false, false}, false},
		{"false2", []bool{false, false, true}, false},
		{"true", []bool{true, true, true}, true},
	}
	for _, tt := range tt {
		t.Run(tt.name, func(t *testing.T) {
			got := itlib.All(sliceit.In(tt.args), func(item bool) bool {
				return item
			})
			assertpkg.Equal(t, tt.wantOk, got)
		})
	}
}
