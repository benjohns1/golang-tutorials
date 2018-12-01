package main

import "fmt"

func main() {
	assignAndCallAnonymous()
	justCallAnonymous()
	anonymousWithArgument()
	fmt.Println("---")
	userDefinedFunctionTypes()
	fmt.Println("---")
	functionAsArgument()
	returnFunction()
	fmt.Println("---")
	closureExample1()
	closureExample2()
	fmt.Println("---")
	studentFilterExample()
}

func assignAndCallAnonymous() {
	a := func() {
		fmt.Println("hello world first class function: define, assign, and call anonymous function")
	}
	a()
	fmt.Printf("Anonymous function type: %T\n", a)
}

func justCallAnonymous() {
	func() {
		fmt.Println("hello world first class function: just define and call anonymous function")
	}()
}

func anonymousWithArgument() {
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")
}

// User defined function types
type add func(a int, b int) int

func userDefinedFunctionTypes() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(5, 6)
	fmt.Println("Sum", s)
}

// Higher-order functions

func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func functionAsArgument() {
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}

func simpleReturn() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func returnFunction() {
	s := simpleReturn()
	fmt.Println(s(60, 7))
}

// Closures

func closureExample1() {
	a := 5
	func() {
		fmt.Println("a =", a)
	}()
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func closureExample2() {
	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}

// Practical example
type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
	age       int
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func modifyAges(s []student, f func(int) int) []student {
	var r []student
	for _, v := range s {
		v.age = f(v.age)
		r = append(r, v)
	}
	return r
}

func studentFilterExample() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	filterByGradeB := func(s student) bool {
		return s.grade == "B"
	}
	filterByCountryIndia := func(s student) bool {
		return s.country == "India"
	}
	addOneToAge := func(age int) int {
		return age + 1
	}
	f1 := filter(s, filterByGradeB)
	f2 := filter(s, filterByCountryIndia)
	f3 := modifyAges(s, addOneToAge)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
}
