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

package itlib

import (
	"github.com/0x5a17ed/itkit"
)

type EmptyIterator[T any] struct{}

// Ensure EmptyIterator implements the iterator interface.
var _ itkit.Iterator[struct{}] = &EmptyIterator[struct{}]{}

func (e EmptyIterator[T]) Next() bool   { return false }
func (e EmptyIterator[T]) Value() (v T) { return }

// Empty returns an Iterator that is always exhausted.
func Empty[T any]() itkit.Iterator[T] { return &EmptyIterator[T]{} }
