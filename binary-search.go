package main

import (
	"fmt"
)

func binarySearch(v int, lst []int, low int, high int) bool {
	if low > high {
		fmt.Println("not found")
		return false
	}

	middle := (low + high) / 2
	if v == lst[middle] {
		fmt.Printf("found! It is at %d\n", middle)
		return true
	} else if v > lst[middle] {
		return binarySearch(v, lst, middle+1, high)
	} else {
		return binarySearch(v, lst, low, middle-1)
	}
}

func main() {
	lst := []int{1, 2, 3, 4, 5, 6, 7}
	binarySearch(5, lst, 0, 6)
}
