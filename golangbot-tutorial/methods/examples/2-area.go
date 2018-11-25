package examples

import (
	"fmt"
	"math"
)

type rectangle struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func Area() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.area())
	c := circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.area())
}
