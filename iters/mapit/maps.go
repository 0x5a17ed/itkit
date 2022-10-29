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

package mapit

import (
	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/genit"
	"github.com/0x5a17ed/itkit/itlib"
	"github.com/0x5a17ed/itkit/ittuple"
)

// Ensure T2 conforms to the Pair protocol.
var _ itlib.Pair[struct{}, struct{}] = &ittuple.T2[struct{}, struct{}]{}

// To consumes an [Iterator] returning its Pair elements as a Go map.
func To[K comparable, V any](it itkit.Iterator[itlib.Pair[K, V]]) (out map[K]V) {
	return itlib.ApplyTo(it, make(map[K]V), func(m map[K]V, p itlib.Pair[K, V]) {
		l, r := p.Values()
		m[l] = r
	})
}

// In returns an [GIterator] yielding Pair items in the given map.
func In[K comparable, V any](m map[K]V) *genit.GIterator[itlib.Pair[K, V]] {
	return genit.Generator(func(g *genit.G[itlib.Pair[K, V]]) {
		for k, v := range m {
			g.Send(ittuple.T2[K, V]{Left: k, Right: v})
		}
	})
}

// Values returns an [GIterator] yielding the values of the given Go map.
func Values[K comparable, V any](m map[K]V) *genit.GIterator[V] {
	return genit.Generator(func(g *genit.G[V]) {
		for _, v := range m {
			g.Send(v)
		}
	})
}

// Keys returns an [GIterator] yielding the keys of the given Go map.
func Keys[K comparable, V any](m map[K]V) *genit.GIterator[K] {
	return genit.Generator(func(g *genit.G[K]) {
		for k := range m {
			g.Send(k)
		}
	})
}
