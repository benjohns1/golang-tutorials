package examples

import (
	"fmt"
)

type employee2 struct {
	name string
	age  int
}

func (e employee2) changeName(newName string) {
	e.name = newName
}

func (e *employee2) changeAge(newAge int) {
	e.age = newAge
}

func PointerReceiver() {
	e := employee2{
		name: "Mark Andrew",
		age:  50,
	}

	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\nEmployee age before change: %d", e.age)
	(&e).changeAge(51) // can be notated as: e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age)
}
