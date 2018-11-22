package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	// Calculate bill
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	// Rectangle
	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	// Rectangle 2
	area2, perimeter2 := rectProps2(2.6, 1.7)
	fmt.Printf("Area2 %f Perimeter2 %f\n", area2, perimeter2)

	// Rectangle blank identifier
	area3, _ := rectProps2(3.7, 1.2)
	fmt.Printf("Area3 %f\n", area3)
}
