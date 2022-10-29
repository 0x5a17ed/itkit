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

package itlib_test

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestChain(t *testing.T) {
	type args struct {
		iters []itkit.Iterator[int]
	}
	tests := []struct {
		name   string
		args   args
		wanted []int
	}{
		{"none", args{[]itkit.Iterator[int]{}}, []int(nil)},
		{"empty", args{[]itkit.Iterator[int]{itlib.Empty[int]()}}, []int(nil)},

		{"one", args{[]itkit.Iterator[int]{
			sliceit.In([]int{1, 2, 3}),
		}}, []int{1, 2, 3}},

		{"two", args{[]itkit.Iterator[int]{
			sliceit.In([]int{1, 2, 3}),
			sliceit.In([]int{7, 8, 9}),
		}}, []int{1, 2, 3, 7, 8, 9}},

		{"empty mixed-start", args{[]itkit.Iterator[int]{
			itlib.Empty[int](),
			sliceit.In([]int{1, 2, 3}),
			sliceit.In([]int{7, 8, 9}),
		}}, []int{1, 2, 3, 7, 8, 9}},
		{"empty mixed-middle", args{[]itkit.Iterator[int]{
			sliceit.In([]int{1, 2, 3}),
			itlib.Empty[int](),
			sliceit.In([]int{7, 8, 9}),
		}}, []int{1, 2, 3, 7, 8, 9}},
		{"empty mixed-end", args{[]itkit.Iterator[int]{
			sliceit.In([]int{1, 2, 3}),
			sliceit.In([]int{7, 8, 9}),
			itlib.Empty[int](),
		}}, []int{1, 2, 3, 7, 8, 9}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sliceit.To(itlib.ChainV(test.args.iters...))
			assertpkg.Equal(t, test.wanted, got)
		})
	}
}
