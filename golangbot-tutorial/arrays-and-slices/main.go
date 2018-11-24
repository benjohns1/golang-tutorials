package main

import (
	"fmt"
)

func main() {
	var a [3]int
	fmt.Println(a)
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{5, 6, 7}
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)

	d := [...]int{2, 3, 4, 5}
	fmt.Println(d)

	fmt.Println("---")
	valueType()

	fmt.Println("---")
	passByValue()

	fmt.Println("---")
	arrayLength()

	fmt.Println("---")
	forLoop()

	fmt.Println("---")
	multiDimensional()

	fmt.Println("---")
	slice()

	fmt.Println("---")
	modifySlice()

	fmt.Println("---")
	sliceLength()

	fmt.Println("---")
	sliceMake()

	fmt.Println("---")
	sliceAppend()

	fmt.Println("---")
	variadicSlice()

	fmt.Println("---")
	copySlice()
}

func copySlice() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries)
	return countriesCpy
}

func variadicSlice() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}

func sliceAppend() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "length", len(cars), "capacity", cap(cars))
	moreCars := append(cars, "Toyota")
	fmt.Println("cars:", cars, "length", len(cars), "capacity", cap(cars))
	fmt.Println("more cars:", moreCars, "length", len(moreCars), "capacity", cap(moreCars))

	var names []string
	if names == nil {
		fmt.Println("slice is nil")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names", names)
	}
}

func sliceMake() {
	i := make([]int, 5, 5)
	fmt.Println(i)
}

func sliceLength() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fmt.Println(fruitarray)
	fruitslice := fruitarray[1:3]
	fmt.Printf("array length %d, slice length %d, slice capacity %d\n", len(fruitarray), len(fruitslice), cap(fruitslice))
	fmt.Println(fruitslice)

	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("after reslicing: slice length", len(fruitslice), ", slice capacity", cap(fruitslice))
	fmt.Println(fruitslice)
}

func modifySlice() {
	darr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:]
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after change 1", numa)
	nums2[1] = 101
	fmt.Println("array after change 2", numa)
}

func slice() {
	a := [5]int{1, 2, 3, 4, 5}
	var b = a[1:4]
	fmt.Println(b)

	var c = []int{6, 7, 8}
	fmt.Println(c)
}

func multiDimensional() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printMultiArray(a)
	var b [3][2]string
	b[0][0] = "Apple"
	b[0][1] = "Samsung"
	b[1][0] = "Microsoft"
	b[1][1] = "Google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printMultiArray(b)
}

func printMultiArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func forLoop() {
	a := [...]float64{67.7, 89.8, 21, 78}
	for i := 0; i < len(a); i++ {
		fmt.Printf("element %d of a is %.2f\n", i, a[i])
	}
	fmt.Println()

	sum := float64(0)
	for i, v := range a {
		fmt.Printf("element %d of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

func arrayLength() {
	a := [...]float64{67.7, 89.8, 21, 78}
	fmt.Println("length of a is", len(a))
}

func passByValue() {
	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changeLocal(num)
	fmt.Println("after passing to function", num)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function", num)
}

func valueType() {
	a := [...]string{"usa", "china", "india", "germany", "france"}
	b := a
	b[0] = "singapore"
	fmt.Println("a is", a)
	fmt.Println("b is", b)
}
