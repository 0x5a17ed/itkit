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

type ApplyFunc[T any] func(item T)

// Apply walks through the given Iterator it and calls ApplyFunc fn
// for every single entry.
func Apply[T any](it Iterator[T], fn ApplyFunc[T]) {
	for it.Next() {
		fn(it.Value())
	}
}

type ApplyNFunc[T any] func(i int, item T)

// ApplyN walks through the given Iterator it and calls ApplyNFunc fn
// for every single entry together with its index.
func ApplyN[T any](it Iterator[T], fn ApplyNFunc[T]) {
	for i := 0; it.Next(); i += 1 {
		fn(i, it.Value())
	}
}

type EachFunc[T any] func(item T) bool

// Each walks through the given Iterator it and calls EachFunc fn for
// every single entry, aborting if EachFunc fn returns true.
func Each[T any](it Iterator[T], fn EachFunc[T]) {
	for it.Next() {
		if fn(it.Value()) {
			break
		}
	}
}

type EachNFunc[T any] func(i int, item T) bool

// EachN walks through the given Iterator it and calls EachNFunc fn for
// every single entry together with its index, aborting if EachFunc fn
// returns true.
func EachN[T any](it Iterator[T], fn EachNFunc[T]) {
	for i := 0; it.Next(); i += 1 {
		if fn(i, it.Value()) {
			break
		}
	}
}
