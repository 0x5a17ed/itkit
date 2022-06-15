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
	"github.com/0x5a17ed/itkit"
)

func main() {
	it := &itkit.SliceIterator[int]{Data: []int{1, 2, 3}}
	for it.Next() {
		fmt.Println(it.Value)()
	}
}
```
