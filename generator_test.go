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
	"runtime"
	"testing"
	"time"

	assertpkg "github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/0x5a17ed/itkit"
)

func TestGenerator(t *testing.T) {
	defer goleak.VerifyNone(t)

	s := itkit.ToSlice(itkit.Generator(func(g *itkit.G[int]) {
		for i := 1; i < 5; i++ {
			g.Send(i)
		}
	}).Iter())

	assertpkg.Equal(t, []int{1, 2, 3, 4}, s)
}

func TestGenerator_StopTwice(t *testing.T) {
	defer goleak.VerifyNone(t)

	completedCh := make(chan struct{})

	go func() {
		defer close(completedCh)

		g := itkit.Generator(func(g *itkit.G[int]) {})
		g.Stop()
		g.Stop()
	}()

	select {
	case <-completedCh:
	case <-time.After(5 * time.Second):
		assertpkg.FailNowf(t, "timeout", "test timed out after %s", 5*time.Second)
	}
}

func TestGenerator_CloseEarly(t *testing.T) {
	defer goleak.VerifyNone(t)

	stoppedAt := make(chan int, 1)

	g := itkit.Generator(func(g *itkit.G[int]) {
		var i int

		defer func() {
			stoppedAt <- i
		}()

		for ; i < 15; i++ {
			g.Send(i)
		}
	})

	itkit.Apply(g.Iter(), func(i int) {
		if i == 3 {
			g.Stop()
		}
	})

	assertpkg.Equal(t, 4, <-stoppedAt)
}

func TestGenerator_IterateHalf(t *testing.T) {
	defer goleak.VerifyNone(t)

	g := itkit.Generator(func(g *itkit.G[int]) {
		for i := 1; i < 10; i++ {
			g.Send(i)
		}
	})

	var s []int
	itkit.Each(g.Iter(), func(i int) bool {
		// Calling GC in the loop trying to get rid of the generator.
		runtime.GC()
		s = append(s, i)
		return i > 4
	})

	// Ensure the generator yielded all the items needed.
	assertpkg.Equal(t, []int{1, 2, 3, 4, 5}, s)

	// Finally get rid of the generator.
	runtime.GC()
}

func TestGenerator_PanicPropagation(t *testing.T) {
	defer goleak.VerifyNone(t)

	assert := assertpkg.New(t)

	var i *int
	g := itkit.Generator(func(g *itkit.G[int]) {
		g.Send(13)
		*i = 42
	})

	// Advancing the iterator the first time will yield the
	// value 13.
	var ok bool
	assert.NotPanics(func() {
		ok = g.Iter().Next()
	})
	assert.True(ok)
	assert.Equal(13, g.Iter().Value())

	// Advancing the iterator again should crash from the null-pointer
	// assignment.
	assert.PanicsWithError("runtime error: invalid memory address or nil pointer dereference", func() {
		g.Iter().Next()
	})
}

func TestGenerator_PanicOtherChannelError(t *testing.T) {
	defer goleak.VerifyNone(t)

	ch := make(chan struct{})
	close(ch)
	g := itkit.Generator(func(g *itkit.G[int]) {
		ch <- struct{}{}
	})

	// Advancing the iterator again should crash from the null-pointer
	// assignment.
	assertpkg.PanicsWithError(t, "send on closed channel", func() {
		g.Iter().Next()
	})
}

func TestGenerator_PanicPropagationOtherValue(t *testing.T) {
	defer goleak.VerifyNone(t)

	ch := make(chan struct{})
	close(ch)
	g := itkit.Generator(func(g *itkit.G[int]) {
		ch <- struct{}{}
	})

	// Advancing the iterator again should crash from the null-pointer
	// assignment.
	assertpkg.PanicsWithError(t, "send on closed channel", func() {
		g.Iter().Next()
	})
}
