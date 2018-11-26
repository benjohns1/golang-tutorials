package main

import (
	"fmt"
)

type tester interface {
	Test()
}

type myFloat float64

func (m myFloat) Test() {
	fmt.Println(m)
}

func describe(t tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t tester
	f := myFloat(89.7)
	t = f
	describe(t)
	t.Test()
}
