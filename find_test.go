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

func TestFind(t *testing.T) {
	compareFn := func(a, b int) bool { return a >= b }

	tt := []struct {
		name    string
		needle  int
		wantOut int
		wantOk  bool
	}{
		// TODO: test cases
		{"ok", 5, 5, true},
		{"not ok", 6, 0, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assertpkg.New(t)

			it := itkit.From([]int{1, 2, 3, 4, 5})

			gotOut, gotOk := itkit.Find(it, compareFn, tc.needle)
			assert.Equal(tc.wantOk, gotOk)
			assert.Equal(tc.wantOut, gotOut)
		})
	}
}
