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

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	type args[T any] struct {
		src itkit.Iterator[T]
		fn  itlib.TakeWhileFn[T]
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		want      []T
		remainder []T
	}
	tt := []testCase[int]{
		{
			name:      "empty",
			args:      args[int]{itlib.Empty[int](), func(x int) bool { return true }},
			want:      nil,
			remainder: nil,
		},
		{
			name:      "none",
			args:      args[int]{rangeit.Range(5), func(x int) bool { return x < 0 }},
			want:      nil,
			remainder: []int{1, 2, 3, 4},
		},
		{
			name:      "some",
			args:      args[int]{rangeit.Range(5), func(x int) bool { return x < 3 }},
			want:      []int{0, 1, 2},
			remainder: []int{4},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			asserter := assert.New(t)

			it := itlib.TakeWhile(tc.args.src, tc.args.fn)

			got := sliceit.To(it)

			asserter.False(it.Next())

			remainder := sliceit.To(tc.args.src)

			asserter.Equalf(tc.want, got, "TakeWhile(...) output")
			asserter.Equalf(tc.remainder, remainder, "TakeWhile(...) remainder")
		})
	}
}
