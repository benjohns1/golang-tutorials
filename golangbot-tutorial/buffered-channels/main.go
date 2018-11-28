package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	example1()
	fmt.Println("---")

	example2()
	fmt.Println("---")

	lengthVsCapacity()
	fmt.Println("---")

	waitGroups()
	fmt.Println("---")
}

func waitGroups() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started goroutine", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("goroutine %d ended\n", i)
	wg.Done()
}

func lengthVsCapacity() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

func example1() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func example2() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(1 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(1 * time.Second)
	}
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
