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

package runeit

import (
	"unicode/utf8"

	"github.com/0x5a17ed/itkit"
)

// StringIterator represents an iterator yielding individual runes
// from a string.
type StringIterator struct {
	value string

	nonASCIIStart int

	bytePos int
	current rune
}

func (it *StringIterator) Value() rune { return it.current }

func (it *StringIterator) Next() bool {
	if it.bytePos >= len(it.value) {
		return false
	}

	r, w := rune(0), 0
	if it.bytePos < it.nonASCIIStart {
		r, w = rune(it.value[it.bytePos]), 1
	} else {
		r, w = utf8.DecodeRuneInString(it.value[it.bytePos:])
	}
	it.current = r
	it.bytePos += w
	return true
}

// InString returns an iterator which yields all runes in the given string.
func InString(v string) itkit.Iterator[rune] {
	it := &StringIterator{value: v}
	for i := 0; i < len(it.value); i++ {
		if it.value[i] >= utf8.RuneSelf {
			it.nonASCIIStart = i
			return it
		}
	}
	it.nonASCIIStart = len(it.value)
	return it
}
