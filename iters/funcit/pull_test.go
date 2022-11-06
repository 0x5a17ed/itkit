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

package funcit

import (
	"testing"

	assertpkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/sliceit"
)

func slicePullFn[T any](s []T) func() (T, bool) {
	i := 0
	return func() (v T, ok bool) {
		if ok = i < len(s); ok {
			v = s[i]
			i++
		}
		return
	}
}

func emptyPullFn() (int, bool) {
	return 0, false
}

func TestPullIterator(t *testing.T) {
	type args struct {
		puller PullFnT[int]
	}
	tt := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{emptyPullFn}, nil},
		{"one", args{slicePullFn([]int{11})}, []int{11}},
		{"multi", args{slicePullFn([]int{11, 13, 17})}, []int{11, 13, 17}},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assertpkg.Equal(t, tc.want, sliceit.To(PullFn(tc.args.puller)))
		})
	}
}
