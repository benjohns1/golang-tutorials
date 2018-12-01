package main

import (
	"fmt"
	"math"
)

func main() {
	formatError()
	fmt.Println()

	typeAndFieldErrorDetail()
	fmt.Println()

	typeAndMethodErrorDetail()
	fmt.Println()
}

type areaErrorDetail struct {
	err    string
	length float64
	width  float64
}

func (e *areaErrorDetail) Error() string {
	return e.err
}

func (e *areaErrorDetail) lengthNegative() bool {
	return e.length < 0
}

func (e *areaErrorDetail) widthNegative() bool {
	return e.width < 0
}

func rectAreaDetail(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaErrorDetail{err, length, width}
	}
	return length * width, nil
}

func typeAndMethodErrorDetail() {
	length, width := -5.0, -9.0
	area, err := rectAreaDetail(length, width)
	if err != nil {
		if err, ok := err.(*areaErrorDetail); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)

			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)

			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("area of rect", area)
}

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func typeAndFieldErrorDetail() {
	radius := -20.0
	area, err := circleArea2(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero: %s", err.radius, err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle %0.2f", area)
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func formatError() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
