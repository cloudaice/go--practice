package main

import (
	"fmt"
)

type I interface {
	foo()
}

type T struct{}

func (t *T) foo() {
	fmt.Println("t")
}

type V struct{}

func (v *V) foo() {
	fmt.Println("v")
}

func main() {
	var x I = &T{}
	x.foo()

	x = &V{}
	x.foo()
}
