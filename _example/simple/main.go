package main

import (
	"fmt"

	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func main() {
	s := []int{1, 2, 3}

	// iterating using the for keyword.
	for it := sliceit.In(s); it.Next(); {
		fmt.Println(it.Value())
	}

	// iterating using a slightly more functional approach.
	itlib.Apply(sliceit.In(s), func(v int) {
		fmt.Println(v)
	})
}
