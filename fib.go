package fib

import (
	"fmt"
)

func Fib(num int) {
	var current, next int = 0, 1

	for i := 0; i < num; i++ {
		temp := current
		current = next
		next = temp + next
		fmt.Println(next)
	}
}

func RecursiveFib(mem *[]int, n int) int {
	if n == 0 {
		return 0
	}

	if (*mem)[n] != 0 {
		return (*mem)[n]
	}

	result := RecursiveFib(mem, n-1) + RecursiveFib(mem, n-2)
	(*mem)[n] = result
	fmt.Println(result)
	return result

}

func main() {
	var mem []int = make([]int, 100)
	mem[0], mem[1] = 0, 1

	var num int = 10
	// Fib(num)
	RecursiveFib(&mem, num)
}
