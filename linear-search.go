package main

import (
	"fmt"
)

func linearSearch(lst []int, v int) {
	n := len(lst)
	for i := 0; i < n; i++ {
		if lst[i] == v {
			fmt.Printf("found the target: %d at index %d\n", v, i)
			return
		}
	}
	fmt.Printf("Target: %d not found\n", v)
	return
}
