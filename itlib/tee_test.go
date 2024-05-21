// Copyright (c) 2024 individual contributors.
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
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/0x5a17ed/itkit/iters/funcit"
	"github.com/0x5a17ed/itkit/iters/rangeit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/0x5a17ed/itkit/ittuple"
	"github.com/stretchr/testify/assert"
)

type refCounter struct {
	finalized chan int
	n         uint32
}

func (c *refCounter) new() *int {
	n := new(int)
	*n = int(atomic.AddUint32(&c.n, 1))

	runtime.SetFinalizer(n, func(r *int) {
		atomic.AddUint32(&c.n, ^uint32(0))
		c.finalized <- *r
	})

	return n
}

func newRefCounter() *refCounter {
	return &refCounter{
		finalized: make(chan int, 1),
	}
}

func TestTee(t *testing.T) {
	t.Run("n", func(t *testing.T) {
		assert.Len(t, itlib.TeeN(itlib.Empty[struct{}](), 0), 0)
		assert.Len(t, itlib.TeeN(itlib.Empty[struct{}](), 1), 1)
		assert.Len(t, itlib.TeeN(itlib.Empty[struct{}](), 12), 12)
	})

	t.Run("empty", func(t *testing.T) {
		asserter := assert.New(t)

		l, r := itlib.Tee(itlib.Empty[int]())

		asserter.False(l.Next())
		asserter.False(r.Next())
	})

	t.Run("interleaved", func(t *testing.T) {
		l, r := itlib.Tee(rangeit.Range(5))
		s := sliceit.To(itlib.Zip(l.Iter(), r.Iter()))

		assert.Equal(t, []itlib.Pair[int, int]{
			ittuple.T2[int, int]{Left: 0, Right: 0},
			ittuple.T2[int, int]{Left: 1, Right: 1},
			ittuple.T2[int, int]{Left: 2, Right: 2},
			ittuple.T2[int, int]{Left: 3, Right: 3},
			ittuple.T2[int, int]{Left: 4, Right: 4},
		}, s)
	})

	t.Run("progressive", func(t *testing.T) {
		n, m := 100, 10

		asserter := assert.New(t)

		expected := sliceit.To(rangeit.Range(n))

		its := itlib.TeeN(rangeit.Range(n), m)
		for i := 0; i < m; i++ {
			actual := sliceit.To(its[i].Iter())
			asserter.Equal(expected, actual)
		}
	})

	t.Run("concurrent", func(t *testing.T) {
		n, m := 1000, 100
		its := itlib.TeeN(rangeit.Range(n), m)

		c := make(chan []int, m)

		var wg sync.WaitGroup
		for i := 0; i < m; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				c <- sliceit.To(its[i].Iter())
			}(i)
		}
		wg.Wait()

		expected := sliceit.To(rangeit.Range(n))
		asserter := assert.New(t)
		for i := 0; i < m; i++ {
			asserter.Equal(expected, <-c)
		}
	})

	t.Run("copy", func(t *testing.T) {
		asserter := assert.New(t)

		l, _ := itlib.Tee(rangeit.Range(5))

		// Advanced the tee iterator once.
		asserter.Equal(0, itlib.HeadOrElse(l.Iter(), -1))

		// Create a copy of the tee iterator.
		c := l.Copy()
		asserter.NotSame(c, l)

		// Assert retrieving the next item from the clone
		// yields the next value.
		asserter.Equal(1, itlib.HeadOrElse(c.Iter(), -1))

		// Assert retrieving the next item from the original
		// tee reader yields the same value.
		asserter.Equal(1, itlib.HeadOrElse(l.Iter(), -1))
	})

	t.Run("leak", func(t *testing.T) {
		// Force garbage collection to start with a clean slate.
		runtime.GC()

		asserter := assert.New(t)

		c := newRefCounter()

		it := funcit.PullFn(func() (*int, bool) {
			return c.new(), true
		})
		it = itlib.Limit(it, 10)

		l, r := itlib.Tee(it)

		// Consume left side first, filling up the buffers.
		asserter.Len(sliceit.To(l.Iter()), 10)

		// Assert that there are currently 10 objects alive.
		asserter.Equal(c.n, uint32(10))

		// Consume half on the right side.
		s := sliceit.To(itlib.Limit(r.Iter(), 5))
		asserter.Len(s, 5)

		// Force garbage collection to clean up references.
		for i := 0; i < 3; i++ {
			runtime.GC()
		}

		// Assert some previously loaded items have
		// been freed up.
		for i := 0; i < 4; i++ {
			select {
			case <-c.finalized:
			case <-time.After(2 * time.Second):
				asserter.FailNow("timeout")
			}
		}
		asserter.Equal(uint32(6), c.n)

		runtime.KeepAlive(r)
	})
}
