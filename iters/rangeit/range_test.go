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

package rangeit_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
)

func TestRange(t *testing.T) {
	type args struct {
		fn func() itkit.Iterator[int]
	}
	tt := []struct {
		name string
		args args
		want []int
	}{
		{"simple", args{fn: func() itkit.Iterator[int] {
			return rangeit.Range(10)
		}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},

		{"In", args{fn: func() itkit.Iterator[int] {
			return rangeit.RangeFrom(5, 10)
		}}, []int{5, 6, 7, 8, 9}},
		{"WithStep", args{fn: func() itkit.Iterator[int] {
			return rangeit.RangeStep(0, 10, 3)
		}}, []int{0, 3, 6, 9}},

		{"negativeStop", args{fn: func() itkit.Iterator[int] {
			return rangeit.Range(-4)
		}}, nil},
		{"negativeStep", args{fn: func() itkit.Iterator[int] {
			return rangeit.RangeStep(5, -5, -3)
		}}, []int{5, 2, -1, -4}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertpkg.New(t)
			assert.Equal(sliceit.To(tc.args.fn()), tc.want)
		})
	}
}
