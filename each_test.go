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

func TestApply(t *testing.T) {
	assert := assertpkg.New(t)

	var loops, output int
	itkit.Apply(itkit.Range(10), func(item int) {
		loops += 1
		output += item
		return
	})

	assert.Equal(10, loops)
	assert.Equal(45, output)
}

func TestApplyN(t *testing.T) {
	assert := assertpkg.New(t)

	var runs, output int
	itkit.ApplyN(itkit.Range(10), func(index, item int) {
		runs += 1
		output += index + item
		return
	})

	assert.Equal(10, runs)
	assert.Equal(90, output)
}

func TestEach(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		assert := assertpkg.New(t)

		var loops, output int
		itkit.Each(itkit.Range(10), func(item int) (o bool) {
			loops += 1
			output += item
			return
		})

		assert.Equal(10, loops)
		assert.Equal(45, output)
	})

	t.Run("half", func(t *testing.T) {
		assert := assertpkg.New(t)

		var loops, output int
		itkit.Each(itkit.Range(10), func(item int) (o bool) {
			loops += 1
			output += item
			return item > 3
		})

		assert.Equal(5, loops)
		assert.Equal(10, output)
	})
}

func TestEachIndex(t *testing.T) {
	t.Run("full", func(t *testing.T) {
		assert := assertpkg.New(t)

		var runs, output int
		itkit.EachN(itkit.Range(10), func(index, item int) (o bool) {
			runs += 1
			output += index + item
			return
		})

		assert.Equal(10, runs)
		assert.Equal(90, output)
	})

	t.Run("half", func(t *testing.T) {
		assert := assertpkg.New(t)

		var runs, output int
		itkit.EachN(itkit.Range(10), func(index, item int) (o bool) {
			runs += 1
			output += index + item
			return index >= 4
		})

		assert.Equal(5, runs)
		assert.Equal(20, output)
	})
}
