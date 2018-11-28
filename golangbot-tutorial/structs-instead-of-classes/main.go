package main

import "golang-tutorials/golangbot-tutorial/structs-instead-of-classes/employee"

func main() {
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
