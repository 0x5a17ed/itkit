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

package rangeit_test

import (
	"testing"

	assertPkg "github.com/stretchr/testify/assert"
	requirePkg "github.com/stretchr/testify/require"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestCount(t *testing.T) {
	type args struct {
		iter itkit.Iterator[int]
	}
	tt := []struct {
		name string
		args args
		want []int
	}{
		{name: "count", args: args{iter: rangeit.Count[int]()},
			want: []int{0, 1, 2, 3, 4}},
		{name: "count-from", args: args{iter: rangeit.CountFrom(10)},
			want: []int{10, 11, 12, 13, 14}},
		{name: "count-step", args: args{iter: rangeit.CountStep(0, 2)},
			want: []int{0, 2, 4, 6, 8}},
		{name: "count-from-step", args: args{iter: rangeit.CountStep(5, 2)},
			want: []int{5, 7, 9, 11, 13}},
		{name: "count-step-negative", args: args{iter: rangeit.CountStep(0, -1)},
			want: []int{0, -1, -2, -3, -4}},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			for i := 0; i < 5; i += 1 {
				v, ok := itlib.Head(tc.args.iter)
				requirePkg.True(t, ok)
				got = append(got, v)
			}
			assertPkg.Equal(t, tc.want, got)
		})
	}
}
