package main

import (
	"fmt"
)

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func basic() {
	findType("Naveen")
	findType(77)
	findType(89.98)
}

func main() {
	basic()
	fmt.Println("---")
	interfaceType()
}

type describer interface {
	describe()
}

type person struct {
	name string
	age  int
}

func (p person) describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case describer:
		v.describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func interfaceType() {
	findType2("Naveen")
	p := person{
		name: "Naveen R",
		age:  25,
	}
	findType2(p)
}
