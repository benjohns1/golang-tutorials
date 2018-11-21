package main

import (
	"fmt"
)

func main() {
	boolTut()
	fmt.Println("")
	complexTut()
	fmt.Println("")
	conversionTut()
}

func boolTut() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)
}

func complexTut() {
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}

func conversionTut() {
	i := 55
	j := 67.8
	sum := i + int(j)
	fmt.Println(sum)
}
