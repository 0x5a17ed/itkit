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

type ApplyFn[T any] func(item T)

// Apply walks through the given Iterator it and calls ApplyFn fn
// for every single entry.
func Apply[T any](it Iterator[T], fn ApplyFn[T]) {
	for it.Next() {
		fn(it.Value())
	}
}

type ApplyNFn[T any] func(i int, item T)

// ApplyN walks through the given Iterator it and calls ApplyNFn fn
// for every single entry together with its index.
func ApplyN[T any](it Iterator[T], fn ApplyNFn[T]) {
	for i := 0; it.Next(); i += 1 {
		fn(i, it.Value())
	}
}

type EachFn[T any] func(item T) bool

// Each walks through the given Iterator it and calls EachFn fn for
// every single entry, aborting if EachFn fn returns true.
func Each[T any](it Iterator[T], fn EachFn[T]) {
	for it.Next() {
		if fn(it.Value()) {
			break
		}
	}
}

type EachNFn[T any] func(i int, item T) bool

// EachN walks through the given Iterator it and calls EachNFn fn for
// every single entry together with its index, aborting if the given
// function returns true.
func EachN[T any](it Iterator[T], fn EachNFn[T]) {
	for i := 0; it.Next(); i += 1 {
		if fn(i, it.Value()) {
			break
		}
	}
}
