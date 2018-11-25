package examples

import (
	"fmt"
)

type rectangle2 struct {
	length int
	width  int
}

func area(r rectangle2) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle2) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func MethodValueReceivers() {
	r := rectangle2{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	   compilation error, cannot use p (type *rectangle2) as type rectangle2
	   in argument to area
	*/
	//area(p)

	p.area() //calling value receiver with a pointer
}
