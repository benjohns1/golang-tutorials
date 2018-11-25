package main

import (
	"fmt"
	"golang-tutorials/golangbot-tutorial/structures/computer"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

type Person struct {
	string
	int
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	emp2 := Employee{"Thomas", "Paul", 29, 800}

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// Anonymous structures
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	// Zero values
	var emp4 Employee
	fmt.Println("Employee 4", emp4)

	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("Employee 5", emp5)

	// Accessing fields
	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Printf("Employee 6 Name: %s %s\n", emp6.firstName, emp6.lastName)

	var emp7 Employee
	emp7.firstName = "Jack"
	emp7.lastName = "Adams"
	fmt.Println("Employee 7:", emp7)

	// Pointers
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	// Anonymous fields
	p := Person{"Naveen", 50}
	fmt.Println(p)

	var p1 Person
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	nested()

	promotedFields()

	exporting()

	equality()
}

func equality() {
	type name struct {
		firstName string
		lastName  string
	}

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	fmt.Println("Equality:", name1 == name2)
}

func exporting() {
	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)
}

func promotedFields() {
	type Address struct {
		city, state string
	}

	type Person struct {
		name    string
		age     int
		Address // <--- !!!
	}

	var p Person
	p.name = "Naveen"
	p.age = 50
	p.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city) // <--- !!!
	fmt.Println("State:", p.Address.state)
}

func nested() {
	type Address struct {
		city, state string
	}

	type Person struct {
		name    string
		age     int
		address Address
	}

	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}
