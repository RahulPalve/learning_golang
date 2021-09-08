package main

import "fmt"

func Bubblesort(arr []int) {
	
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-2; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func main() {
	var array = []int{3, 5, 1, 2, 7, 8}
	fmt.Println(array)

	Bubblesort(array)

	fmt.Println(array)

}
