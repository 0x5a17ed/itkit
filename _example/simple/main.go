package main

import (
	"fmt"

	"github.com/0x5a17ed/itkit"
)

func main() {
	s := []int{1, 2, 3}

	// iterating using the for keyword.
	for it := itkit.InSlice(s); it.Next(); {
		fmt.Println(it.Value())
	}

	// iterating using a slightly more functional approach.
	itkit.Apply(itkit.InSlice(s), func(v int) {
		fmt.Println(v)
	})
}
