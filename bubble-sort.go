package main

import (
	"fmt"
)

func bubbleSort(lst []int) {
	n := len(lst)
	swapped := true
	for swapped == true {
		swapped = false
		for i := 0; i < n-1; i++ {
			if lst[i] > lst[i+1] {
				lst[i], lst[i+1] = lst[i+1], lst[i]
				swapped = true
			}
		}
		fmt.Println(lst)
	}
	fmt.Println(lst)
}

// func main() {
// 	lst := []int{3, 1, 4, 2}
// 	// selectionSort(lst)
// 	bubbleSort(lst)
// 	var a, b, c int
// 	a, b, c = 1, 2, 3
// 	fmt.Println(a, b, c)
// }
