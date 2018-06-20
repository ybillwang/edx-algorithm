package main

import (
	"fmt"
)

func selectionSort(lst []int) {
	n := len(lst)
	for i := 0; i < n; i++ {
		index := 0
		smallest := lst[i]
		for j := i; j < n; j++ {
			if lst[j] <= smallest {
				smallest = lst[j]
				index = j
			}
		}
		temp := lst[i]
		lst[i] = smallest
		lst[index] = temp
	}
	fmt.Println(lst)
}
