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

### Stack

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

### Set

```
package main

import (
	"fmt"
	"github.com/paralleltree/go-collections/collections"
)

func main() {
	s1 := collections.NewSet[int]()
	s1.Add(10)
	s1.Add(20)
	s2 := collections.NewSet[int]()
	s2.Add(10)

	union := s1.Union(s2)
	fmt.Println("union.Contains(10):", union.Contains(10))
	fmt.Println("union.Contains(20):", union.Contains(20))

	intersect := s1.Intersect(s2)
	fmt.Println("intersect.Contains(10):", intersect.Contains(10))
	fmt.Println("intersect.Contains(20):", intersect.Contains(20))

	except := s1.Except(s2)
	fmt.Println("except.Contains(10):", except.Contains(10))
	fmt.Println("except.Contains(20):", except.Contains(20))
}

// Output:
// union.Contains(10): true
// union.Contains(20): true
// intersect.Contains(10): true
// intersect.Contains(20): false
// except.Contains(10): false
// except.Contains(20): true
```
