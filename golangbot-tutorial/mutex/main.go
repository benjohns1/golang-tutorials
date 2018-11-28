package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	mutex()

	channel()
}

func channel() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func mutex() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}
