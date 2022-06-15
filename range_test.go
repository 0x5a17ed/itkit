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
	"testing"

	assertpkg "github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	type args struct {
		stop int
		opts []OptionFunc
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"simple", args{stop: 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"WithStart", args{stop: 10, opts: []OptionFunc{WithStart(5)}}, []int{5, 6, 7, 8, 9}},
		{"WithStep", args{stop: 10, opts: []OptionFunc{WithStep(3)}}, []int{0, 3, 6, 9}},
		{"WithStart,WithStep", args{stop: 10, opts: []OptionFunc{WithStart(1), WithStep(3)}}, []int{1, 4, 7}},

		{"negativeStop", args{stop: -4}, nil},
		{"negativeStep", args{stop: -5, opts: []OptionFunc{WithStart(5), WithStep(-3)}}, []int{5, 2, -1, -4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assertpkg.New(t)
			assert.Equal(Slice(Range(tt.args.stop, tt.args.opts...)), tt.want)
		})
	}
}
