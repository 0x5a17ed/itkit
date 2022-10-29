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

package itkit

type Pair[T1, T2 any] interface{ Values() (T1, T2) }

// ToMap consumes an [Iterator] returning its Pair elements as a Go map.
func ToMap[K comparable, V any](it Iterator[Pair[K, V]]) (out map[K]V) {
	return ApplyTo(it, make(map[K]V), func(m map[K]V, p Pair[K, V]) {
		l, r := p.Values()
		m[l] = r
	})
}

// InMap returns an [GIterator] yielding Pair items in the given map.
func InMap[K comparable, V any](m map[K]V) *GIterator[Pair[K, V]] {
	return Generator(func(g *G[Pair[K, V]]) {
		for k, v := range m {
			g.Send(Tuple2[K, V]{Left: k, Right: v})
		}
	})
}

// Values returns an [GIterator] yielding the values of the given Go map.
func Values[K comparable, V any](m map[K]V) *GIterator[V] {
	return Generator(func(g *G[V]) {
		for _, v := range m {
			g.Send(v)
		}
	})
}

// Keys returns an [GIterator] yielding the keys of the given Go map.
func Keys[K comparable, V any](m map[K]V) *GIterator[K] {
	return Generator(func(g *G[K]) {
		for k := range m {
			g.Send(k)
		}
	})
}
