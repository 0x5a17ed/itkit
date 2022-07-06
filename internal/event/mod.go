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

package event

import (
	"sync"
)

type Signal struct{}

// closedCh represents a closed channel.
var closedCh = make(chan Signal)

func init() {
	close(closedCh)
}

// E is a communication utility for event notification.
type E struct {
	mu sync.Mutex
	ch chan Signal
	is bool
}

// Wait returns a channel for waiting for the event.
func (e *E) Wait() <-chan Signal {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.ch == nil {
		e.ch = make(chan Signal)
	}
	return e.ch
}

// Set marks this event as 'done' and wakes up all waiters.
func (e *E) Set() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !e.is {
		if e.ch == nil {
			e.ch = closedCh
		} else {
			close(e.ch)
		}
		e.is = true
	}
}
