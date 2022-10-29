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

	"github.com/0x5a17ed/itkit/iters/runeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestMap(t *testing.T) {
	assert := assertpkg.New(t)
	s := sliceit.To(itlib.Map(runeit.InString("Hello World!"), func(r rune) string { return string(r) }))
	assert.Equal([]string{"H", "e", "l", "l", "o", " ", "W", "o", "r", "l", "d", "!"}, s)
}
