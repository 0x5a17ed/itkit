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

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestFind(t *testing.T) {
	compareFn := func(a, b int) bool { return a >= b }

	tt := []struct {
		name    string
		needle  int
		wantOut int
		wantOk  bool
	}{
		{"ok", 5, 5, true},
		{"not ok", 6, 0, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertpkg.New(t)

			it := sliceit.In([]int{1, 2, 3, 4, 5})

			gotOut, gotOk := itlib.Find(it, compareFn, tc.needle)
			assert.Equal(tc.wantOk, gotOk)
			assert.Equal(tc.wantOut, gotOut)
		})
	}
}

func TestHead(t *testing.T) {
	type args struct {
		it itkit.Iterator[int]
	}
	type want struct {
		out int
		ok  bool
	}
	tt := []struct {
		name string
		args args
		want []want
	}{
		{"empty", args{it: itlib.Empty[int]()}, []want{{0, false}}},
		{"one", args{it: sliceit.In([]int{23})}, []want{{23, true}, {0, false}}},
		{"multi", args{it: sliceit.In([]int{23, 29})}, []want{{23, true}, {29, true}, {0, false}}},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			for i, w := range tc.want {
				gotOut, gotOk := itlib.Head(tc.args.it)
				assertpkg.Equalf(t, w.out, gotOut, "%d: Head(%v)", i, tc.args.it)
				assertpkg.Equalf(t, w.ok, gotOk, "%d: Head(%v)", i, tc.args.it)
			}
		})
	}
}

func TestHeadOrElse(t *testing.T) {
	t.Run("empty-default", func(t *testing.T) {
		got := itlib.HeadOrElse(itlib.Empty[int](), 23)
		assertpkg.Equal(t, 23, got)
	})

	t.Run("non-empty-not-default", func(t *testing.T) {
		got := itlib.HeadOrElse(sliceit.In([]int{17}), 23)
		assertpkg.Equal(t, 17, got)
	})
}
