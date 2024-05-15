// Copyright (c) 2022 individual contributors.
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

// To builds a Go map from an iterator.
//
// To consumes an [itkit.Iterator] that yields [itlib.Pair] values
// and builds a Go map from all pairs.
func To[K comparable, V any](it itkit.Iterator[itlib.Pair[K, V]]) (out map[K]V) {
	return itlib.ApplyTo(it, make(map[K]V), func(m map[K]V, p itlib.Pair[K, V]) {
		l, r := p.Values()
		m[l] = r
	})
}

// In provides an iterator that yields all keys and values in a Go map.
//
// In returns a [genit.Generator] yielding [itlib.Pair] values which
// reflect all values in the given Go map.
//
// The returned generator will not be automatically stopped.
func In[K comparable, V any](m map[K]V) *genit.Generator[itlib.Pair[K, V]] {
	return genit.Run(func(yield func(itlib.Pair[K, V])) {
		for k, v := range m {
			yield(ittuple.T2[K, V]{Left: k, Right: v})
		}
	})
}

// Values returns a [genit.Generator] yielding all values of the given Go map.
//
// The returned generator will not be automatically stopped.
func Values[K comparable, V any](m map[K]V) *genit.Generator[V] {
	return genit.Run(func(yield func(V)) {
		for _, v := range m {
			yield(v)
		}
	})
}

// Keys returns a [genit.Generator] yielding all keys of the given Go map.
//
// The returned generator will not be automatically stopped.
func Keys[K comparable, V any](m map[K]V) *genit.Generator[K] {
	return genit.Run(func(yield func(K)) {
		for k := range m {
			yield(k)
		}
	})
}
