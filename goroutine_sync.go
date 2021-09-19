package main

import (
	"fmt"
	"sync"
)

func some_thread(a int, b int, total *int, wg *sync.WaitGroup) {

	(*total) = a + b
	wg.Done() // Decrement counter

}

func main() {
	//var total *int = 5 This wont work, memoery is not allocated anywhere remember pntr store memory addresses
	var total *int = new(int)
	(*total) = 5
	var wg sync.WaitGroup
	wg.Add(1) // increament counter
	go some_thread(1, 2, total, &wg)
	wg.Wait() // wait till counter is zero, all goroutines executed
	fmt.Println("Total is : ", *total)
	// Will print 3
	// Will print 5, 3 random rsult 5 in most caese as goroutine haventes finished its exevution yet
	// Check goroutine_sync.go for synchronizzed
}
