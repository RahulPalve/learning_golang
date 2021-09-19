package main

import "fmt"

func getAintoB(ab chan int) {
	ab <- 3 * 2
}

func getCintoD(cd chan int) {
	cd <- 5 - 2

}

func main() {
	ab := make(chan int)
	cd := make(chan int)
	go getAintoB(ab)
	go getCintoD(cd)

	fmt.Println(<-ab * <-cd)

}

// If you observce this when we are calculatinh
// getAintoB we are also calc getCintoD! Multithreading cool! :)
// we are waiting
