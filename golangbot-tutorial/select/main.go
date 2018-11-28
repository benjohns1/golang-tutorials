package main

import (
	"fmt"
	"time"
)

func main() {
	basicSelect()
	fmt.Println("---")

	defaultCase()
	fmt.Println("---")

	defaultAndDeadlock()
	fmt.Println("---")

	randomSelection()
}

func randomSelection() {
	output1 := make(chan string)
	output2 := make(chan string)
	go serverA(output1)
	go serverB(output2)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func serverA(ch chan string) {
	ch <- "from serverA"
}
func serverB(ch chan string) {
	ch <- "from serverB"
}

func defaultAndDeadlock() {
	// no write to channel
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}

	// nil channel
	var ch2 chan string
	select {
	case <-ch2:
	default:
		fmt.Println("default case executed")
	}
}

func defaultCase() {
	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value:", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}

func process(ch chan<- string) {
	time.Sleep(4050 * time.Millisecond)
	ch <- "process successful"
}

func basicSelect() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}
