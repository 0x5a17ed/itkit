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

import (
	"fmt"
)

// Tuple2 represents a generic tuple holding 2 values.
type Tuple2[T1, T2 any] struct {
	Left  T1
	Right T2
}

// Ensure Tuple2 conforms to the Pair protocol.
var _ Pair[struct{}, struct{}] = &Tuple2[struct{}, struct{}]{}

// Len returns the number of values held by the tuple.
func (t Tuple2[T1, T2]) Len() int {
	return 2
}

// Values returns the values held by the tuple.
func (t Tuple2[T1, T2]) Values() (T1, T2) {
	return t.Left, t.Right
}

// Array returns an array of the tuple values.
func (t Tuple2[T1, T2]) Array() [2]any {
	return [2]any{t.Left, t.Right}
}

// Slice returns a slice of the tuple values.
func (t Tuple2[T1, T2]) Slice() []any {
	a := t.Array()
	return a[:]
}

// String returns the string representation of the tuple.
func (t Tuple2[T1, T2]) String() string {
	return fmt.Sprintf("[%#v %#v]", t.Slice()...)
}

func NewTuple2[T1, T2 any](left T1, right T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{Left: left, Right: right}
}
