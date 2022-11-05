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

package chanit

import (
	"github.com/0x5a17ed/itkit"
)

type ChannelIterator[T any] struct {
	ch <-chan T
	v  T
}

func (it *ChannelIterator[T]) Value() T        { return it.v }
func (it *ChannelIterator[T]) Next() (ok bool) { it.v, ok = <-it.ch; return }

func Channel[T any](ch <-chan T) itkit.Iterator[T] {
	return &ChannelIterator[T]{ch: ch}
}
