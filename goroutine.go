package main

import (
	"fmt"
	// "sync"
)

func some_thread(a int, b int, total *int) {

	(*total) = a + b

}

func main() {
	//var total *int = 5 This wont work, memoery is not allocated anywhere remember pntr store memory addresses
	var total *int = new(int)
	(*total) = 5
	go some_thread(1, 2, total)

	fmt.Println("Total is : ", *total)
	// Will print 5, 3 random rsult 5 in most caese as goroutine haventes finished its exevution yet
	// Check goroutine_sync.go for synchronizzed
}
