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

package itlib

import (
	"sync"

	"github.com/0x5a17ed/itkit"
)

type teeNode[T any] struct {
	data T
	next *teeNode[T]
}

type teeState[T any] struct {
	mx  sync.RWMutex
	src itkit.Iterator[T]
}

func (st *teeState[T]) nextLocked(cur *teeNode[T]) (*teeNode[T], bool) {
	return cur.next, cur.next != nil
}

func (st *teeState[T]) nextRW(cur *teeNode[T]) (n *teeNode[T], ok bool) {
	st.mx.Lock()
	defer st.mx.Unlock()

	if n, ok = st.nextLocked(cur); ok {
		return
	}

	if st.src.Next() {
		cur.next = &teeNode[T]{data: st.src.Value()}
	}
	return cur.next, cur.next != nil
}

func (st *teeState[T]) nextR(cur *teeNode[T]) (n *teeNode[T], ok bool) {
	st.mx.RLock()
	defer st.mx.RUnlock()

	return st.nextLocked(cur)
}

func (st *teeState[T]) next(cur *teeNode[T]) (n *teeNode[T], ok bool) {
	if n, ok = st.nextR(cur); ok {
		return
	}
	return st.nextRW(cur)
}

// TeeIterator is an iterator yielding the same items as their
// siblings created from a given source iterator.
//
// A [TeeIterator] can be copied at any time to create more [TeeIterator]
// instances to the same underlying iterator.  Once copied, the two
// [TeeIterator] instances are identical, but separate: they will yield
// the same items but advance individually.  The same item retrieved
// from one [TeeIterator] instance will also be yielded by its sibling
// [TeeIterator] instances.
//
// Once a [TeeIterator] has been created, the source iterator should
// not be used elsewhere.  Advancing the source iterator elsewhere
// will prevent the [TeeIterator] from seeing items retrieved
// elsewhere.  Use copies of the [TeeIterator] instead.
//
// All [TeeIterator] instances are safe to use in goroutines.
type TeeIterator[T any] struct {
	st  *teeState[T]
	cur *teeNode[T]
}

// Ensure TeeIterator conforms to the Iterator protocol.
var _ itkit.Iterator[struct{}] = &TeeIterator[struct{}]{}

// Next implements the [itkit.Iterator.Next] interface.
func (it *TeeIterator[T]) Next() bool {
	if n, ok := it.st.next(it.cur); ok {
		it.cur = n
		return true
	}

	return false
}

// Value implements the [itkit.Iterator.Value] interface.
func (it *TeeIterator[T]) Value() T {
	return it.cur.data
}

// Iter returns the [TeeIterator] as an [itkit.Iterator] value.
func (it *TeeIterator[T]) Iter() itkit.Iterator[T] {
	return it
}

// Copy copies the [TeeIterator] instance.
func (it *TeeIterator[T]) Copy() *TeeIterator[T] {
	c := new(TeeIterator[T])
	*c = *it
	return c
}

func newTee[T any](src itkit.Iterator[T]) *TeeIterator[T] {
	return &TeeIterator[T]{
		st: &teeState[T]{
			src: src,
		},
		cur: &teeNode[T]{},
	}
}

// TeeN returns n new [TeeIterator] values.
func TeeN[T any](src itkit.Iterator[T], n int) []*TeeIterator[T] {
	if n == 0 {
		return nil
	}

	its := make([]*TeeIterator[T], n)
	its[0] = newTee(src)
	for i := 1; i < n; i++ {
		its[i] = its[0].Copy()
	}
	return its
}

// Tee returns 2 new [TeeIterator] values.
func Tee[T any](src itkit.Iterator[T]) (*TeeIterator[T], *TeeIterator[T]) {
	t := newTee(src)
	return t, t.Copy()
}
