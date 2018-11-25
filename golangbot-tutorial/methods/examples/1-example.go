package examples

import (
	"fmt"
)

type employee struct {
	name     string
	salary   int
	currency string
}

func (e employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func Example() {
	emp1 := employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary()
}
