# go-collections

[![Test package](https://github.com/paralleltree/go-collections/actions/workflows/test.yml/badge.svg)](https://github.com/paralleltree/go-collections/actions/workflows/test.yml)

A set of collections with generics in Go.

This package is experimental because the generics in go has not yet completely released.

## Build and Test

You have to work with Go 1.18 to use this package because currently released version(1.17) cannot export generic functions.

You can get latest go version(unstable) that supports exporting generic functions with gotip command and test this package as below:

    $ go version
    go version go1.17.2 linux/amd64
    $ go install golang.org/dl/gotip@latest
    $ gotip download # build and install current master branch
    $ gotip version
    go version devel go1.18-ed1c8db Fri Oct 15 21:46:06 2021 +0000 linux/amd64
    $ gotip test ./...

## Examples

```
package main

import (
	"fmt"
	"github.com/paralleltree/go-collections/collections"
)

func main() {
	s := collections.Stack[int]{}
	s.Push(10)
	s.Push(20)

	for {
		if v, ok := s.Pop(); ok {
			fmt.Println(v)
			continue
		}
		break
	}
}

// Output:
// 20
// 10
```
