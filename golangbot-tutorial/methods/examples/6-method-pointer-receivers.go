package examples

import (
	"fmt"
)

type rectangle3 struct {
	length int
	width  int
}

func perimeter(r *rectangle3) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle3) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func MethodPointerReceivers() {
	r := rectangle3{
		length: 10,
		width:  5,
	}
	p := &r //pointer to r
	perimeter(p)
	p.perimeter()

	/*
	   cannot use r (type rectangle3) as type *rectangle3 in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() //calling pointer receiver with a value

}
