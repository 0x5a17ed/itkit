# itkit

[![License: APACHE-2.0](https://img.shields.io/badge/license-APACHE--2.0-blue?style=flat-square)](https://www.apache.org/licenses/)

Short, dead simple and concise generic iterator interface. With a few extras similar to what python has to offer.


## 📦 Installation

```shell
$ go get -u github.com/0x5a17ed/itkit@latest
```


## 🤔 Usage

```go
package main

import (
	"fmt"

	"github.com/0x5a17ed/itkit"
)

func main() {
	s := []int{1, 2, 3}

	// iterating using the for keyword.
	for it := itkit.From(s); it.Next(); {
		fmt.Println(it.Value())
	}

	// iterating using a slightly more functional approach.
	itkit.Apply(itkit.From(s), func(v int) {
		fmt.Println(v)
	})
}
```


## 🥇 Acknowledgments

The iterator interface is desgined after the stateful iterators pattern explained in the brilliant blog post from <https://ewencp.org/blog/golang-iterators/index.html>. Most functions to manipulate iterators draw inspiration from different sources such as Python and [github.com/samber/lo](https://github.com/samber/lo).
