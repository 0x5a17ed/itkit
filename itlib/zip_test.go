// Copyright (c) 2022 individual contributors
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

	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/0x5a17ed/itkit/ittuple"
)

func TestZip(t *testing.T) {
	type args struct {
		left  []string
		right []int
	}
	tt := []struct {
		name string
		args args
		want []ittuple.T2[string, int]
	}{
		{"empty", args{}, nil},
		{"equal", args{
			left:  []string{"a", "b", "c"},
			right: []int{17, 19, 23},
		}, []ittuple.T2[string, int]{{"a", 17}, {"b", 19}, {"c", 23}}},
		{"left short", args{
			left:  []string{"a", "b"},
			right: []int{17, 19, 23},
		}, []ittuple.T2[string, int]{{"a", 17}, {"b", 19}}},
		{"right short", args{
			left:  []string{"a", "b", "c"},
			right: []int{17, 19},
		}, []ittuple.T2[string, int]{{"a", 17}, {"b", 19}}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := sliceit.To(itlib.Zip(sliceit.In(tc.args.left), sliceit.In(tc.args.right)))

			assertpkg.Equal(t, tc.want, s)
		})
	}
}
