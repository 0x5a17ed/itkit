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

package itlib

import (
	"unicode/utf8"

	"github.com/0x5a17ed/itkit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
)

type StringIter struct {
	value string

	nonASCIIStart int

	bytePos int
	current *rune
}

func (s *StringIter) init() *StringIter {
	for i := 0; i < len(s.value); i++ {
		if s.value[i] >= utf8.RuneSelf {
			s.nonASCIIStart = i
			return s
		}
	}
	s.nonASCIIStart = len(s.value)
	return s
}

func (s *StringIter) Next() bool {
	if s.bytePos >= len(s.value) {
		s.current = nil
		return false
	}

	r, w := rune(0), 0
	if s.bytePos < s.nonASCIIStart {
		r, w = rune(s.value[s.bytePos]), 1
	} else {
		r, w = utf8.DecodeRuneInString(s.value[s.bytePos:])
	}
	s.current = &r
	s.bytePos += w
	return true
}

func (s *StringIter) Value() rune {
	return *s.current
}

// Runes returns an iterator which yields all runes in the given string.
func Runes(v string) itkit.Iterator[rune] {
	return (&StringIter{value: v}).init()
}

// String consumes the given rune iterator and returns the content as
// a string.
func String(it itkit.Iterator[rune]) string {
	return string(sliceit.To(it))
}
