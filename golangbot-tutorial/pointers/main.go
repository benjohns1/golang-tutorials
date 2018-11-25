package main

import "fmt"

func main() {
	pointer()
	fmt.Println("---")

	zeroValue()
	fmt.Println("---")

	dereference()
	fmt.Println("---")

	functionParameter()
	fmt.Println("---")
}

func functionParameter() {
	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function all is", a)
}

func change(val *int) {
	*val = 55
}

func dereference() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
	*a++
	fmt.Println("new value of b is", b)
}

func zeroValue() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after init is", b)
	}
}

func pointer() {
	b := 255
	var a *int = &b
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of b is", a)
}
