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

package runeit_test

import (
	"testing"

	assertPkg "github.com/stretchr/testify/assert"

	"github.com/0x5a17ed/itkit/iters/runeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
)

func TestToString(t *testing.T) {
	type args struct {
		r []rune
	}
	tt := []struct {
		name string
		args args
		want string
	}{
		{"empty-nil", args{r: nil}, ""},
		{"empty-slice", args{r: []rune{}}, ""},

		{"ascii", args{r: []rune{0x61, 0x62, 0x63}}, "abc"},
		{"unicode-1", args{r: []rune{0x65E5, 0x672C}}, "日本"},
		{"unicode-2", args{r: []rune{0x1fae0}}, "\U0001FAE0"}, // 🫠
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := runeit.ToString(sliceit.In(tc.args.r))
			assertPkg.Equalf(t, tc.want, got, "ToString(%v)", tc.args.r)
		})
	}
}
