// Copyright (c) 2023 individual contributors.
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

package genit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/0x5a17ed/itkit/iters/genit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func TestGenerator(t *testing.T) {
	defer goleak.VerifyNone(t)

	g := genit.Run(func(yield func(int)) {
		for i := 1; i < 5; i++ {
			yield(i)
		}
	})
	defer g.Stop()

	s := sliceit.To(g.Iter())

	assert.Equal(t, []int{1, 2, 3, 4}, s)
}

func TestGenerator_IterateHalf(t *testing.T) {
	defer goleak.VerifyNone(t)

	g := genit.Run(func(yield func(int)) {
		for i := 1; ; i++ {
			yield(i)
		}
	})
	defer g.Stop()

	var s []int
	itlib.Each(g.Iter(), func(i int) bool {
		s = append(s, i)
		return i > 4
	})

	// Ensure the generator yielded all the items needed.
	assert.Equal(t, []int{1, 2, 3, 4, 5}, s)
}
