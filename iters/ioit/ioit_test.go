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

package ioit_test

import (
	"io/fs"
	"sync"
	"testing"

	"github.com/0x5a17ed/itkit/iters/ioit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

type Log struct {
	mx      sync.Mutex
	Entries []string
}

func (l *Log) Add(s string) {
	l.mx.Lock()
	defer l.mx.Unlock()

	l.Entries = append(l.Entries, s)
}

func (l *Log) Assert(t *testing.T, s []string) {
	assert.Equal(t, s, l.Entries)
}

func singleGeneratorFactory(l *Log) *ioit.Generator[int] {
	l.Add("factory creating generator")
	gen := ioit.Run(func(cont bool, yield func(int) bool) error {
		l.Add("generator enter")
		defer l.Add("generator leave")

		if !cont {
			l.Add("generator closing down")
			return nil
		}

		l.Add("generator yielding")
		cont = yield(1)
		l.Add("generator yielded")

		if !cont {
			l.Add("generator closing down")
			return fs.ErrClosed
		}

		l.Add("generator returning")
		return fs.ErrNotExist
	})
	l.Add("factory returning generator")

	return gen
}

func badGeneratorFactory(l *Log) *ioit.Generator[int] {
	l.Add("factory creating generator")
	gen := ioit.Run(func(cont bool, yield func(int) bool) error {
		l.Add("generator enter")
		defer l.Add("generator leave")

		for i := 1; ; i++ {
			l.Add("generator yielding")
			yield(i)
			l.Add("generator yielded")
		}
	})
	l.Add("factory returning generator")

	return gen
}

func panicGeneratorFactory(l *Log) *ioit.Generator[int] {
	l.Add("factory creating generator")
	gen := ioit.Run(func(cont bool, yield func(int) bool) error {
		l.Add("generator enter")
		defer l.Add("generator leave")

		l.Add("generator panic")
		panic(nil)
	})
	l.Add("factory returning generator")

	return gen
}

func TestGenerator_Err(t *testing.T) {
	defer goleak.VerifyNone(t)

	asserter := assert.New(t)

	g := ioit.Run(func(cont bool, _ func(any) bool) error {
		if !cont {
			return nil
		}
		return fs.ErrNotExist
	})

	asserter.NoError(g.Err())

	asserter.False(g.Next())

	asserter.ErrorIs(g.Err(), fs.ErrNotExist)
}

func TestGenerator_Close(t *testing.T) {
	t.Run("after completion", func(t *testing.T) {
		asserter := assert.New(t)

		var l Log
		g := singleGeneratorFactory(&l)

		// Exhaust the iterator.
		asserter.Equal([]int{1}, sliceit.To(g.Iter()))

		// Signal the generator to close down.
		l.Add("caller closing")
		asserter.ErrorIs(g.Close(), fs.ErrNotExist)

		l.Assert(t, []string{
			"factory creating generator",
			"factory returning generator",
			"generator enter",
			"generator yielding",
			"generator yielded",
			"generator returning",
			"generator leave",
			"caller closing",
		})
	})

	t.Run("before start", func(t *testing.T) {
		asserter := assert.New(t)

		var l Log
		g := singleGeneratorFactory(&l)

		// Signal the generator to close down.
		l.Add("caller closing")
		asserter.ErrorIs(g.Close(), nil)
		l.Add("caller closed")

		// Assert the generator has reported no error.
		asserter.NoError(g.Err())

		// assert the state at the end.
		l.Assert(t, []string{
			"factory creating generator",
			"factory returning generator",
			"caller closing",
			"generator enter",
			"generator closing down",
			"generator leave",
			"caller closed",
		})
	})

	t.Run("running", func(t *testing.T) {
		asserter := assert.New(t)

		var l Log
		g := singleGeneratorFactory(&l)

		// Retrieve first value from the generator.
		l.Add("caller resuming")
		ok := g.Next()
		l.Add("caller resumed")
		asserter.True(ok)

		// Signal the generator to close down.
		l.Add("caller closing")
		asserter.ErrorIs(g.Close(), fs.ErrClosed)
		l.Add("caller closed")

		// assert the state at the end.
		l.Assert(t, []string{
			"factory creating generator",
			"factory returning generator",
			"caller resuming",
			"generator enter",
			"generator yielding",
			"caller resumed",
			"caller closing",
			"generator yielded",
			"generator closing down",
			"generator leave",
			"caller closed",
		})
	})

	t.Run("bogus", func(t *testing.T) {
		asserter := assert.New(t)

		var l Log
		g := badGeneratorFactory(&l)

		// Signal the generator to close down.
		l.Add("caller closing")
		asserter.ErrorIs(g.Close(), ioit.ErrStopped)
		l.Add("caller closed")

		l.Assert(t, []string{
			"factory creating generator",
			"factory returning generator",
			"caller closing",
			"generator enter",
			"generator yielding",
			"generator leave",
			"caller closed",
		})
	})

	t.Run("panic", func(t *testing.T) {
		asserter := assert.New(t)

		var l Log
		g := panicGeneratorFactory(&l)

		// Signal the generator to close down.
		l.Add("caller closing")
		asserter.Panics(func() {
			_ = g.Close()
		})
		l.Add("caller closed")

		l.Assert(t, []string{
			"factory creating generator",
			"factory returning generator",
			"caller closing",
			"generator enter",
			"generator panic",
			"generator leave",
			"caller closed",
		})
	})
}
