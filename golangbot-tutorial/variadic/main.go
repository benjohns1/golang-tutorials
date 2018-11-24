package main

import (
	"fmt"
)

func main() {

	gotcha()
}

func gotcha() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change(s ...string) {
	s[0] = "go"
	s = append(s, "playground")
	fmt.Println(s)
}
