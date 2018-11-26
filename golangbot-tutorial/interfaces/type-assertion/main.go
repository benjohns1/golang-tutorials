package main

import (
	"fmt"
)

func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)
}

func safeAssert(i interface{}) {
	v, ok := i.(int)
	if !ok {
		fmt.Println("not an integer")
		return
	}
	fmt.Println(v)
}

func main() {
	var intValue interface{} = 56
	assert(intValue)

	var stringValue interface{} = "asdf"
	//assert(stringValue) // runtime error "panic: interface conversion: interface {} is string, not int"
	safeAssert(stringValue)
}
