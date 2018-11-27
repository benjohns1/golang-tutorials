package main

import (
	"fmt"
	"golang-tutorials/golangbot-tutorial/channels/calc"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	time.Sleep(1 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	fmt.Println("---")

	done := make(chan bool)
	go hello(done)
	<-done
	fmt.Println("main function")
	fmt.Println("---")

	// Multiple channels
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
	fmt.Println("---")

	// Unidirectional channels
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
	fmt.Println("---")

	// Closing channels
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received", v, ok)
	}
	fmt.Println("---")

	// For range with channel
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received", v)
	}
	fmt.Println("---")

	calc.Run()
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}
