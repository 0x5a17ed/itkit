package main

import (
	"fmt"

	"github.com/0x5a17ed/itkit/iters/mapit"
	"github.com/0x5a17ed/itkit/iters/sliceit"
	"github.com/0x5a17ed/itkit/itlib"
)

func main() {
	numbers := []int{1, 2, 3}
	strings := []string{"one", "two", "three"}

	m := mapit.To(itlib.Zip(sliceit.In(strings), sliceit.In(numbers)))

	fmt.Println(m)
}
