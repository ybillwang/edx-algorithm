package main

func mergeSort(lst []int) []int {
	n := len(lst)
	var left []int
	var right []int
	for i := 0; i < n; i++ {
		if i < n/2 {
			left = append(left, lst[i])
		} else {
			right = append(right, lst[i])
		}
	}
	if len(left) > 1 {
		left = mergeSort(left)
	}
	if len(right) > 1 {
		right = mergeSort(right)
	}

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[index] = left[i]
			i++
		} else {
			result[index] = right[j]
			j++
		}
		index++
	}
	for i < len(left) {
		result[index] = left[i]
		index++
		i++
	}
	for j < len(right) {
		result[index] = right[j]
		index++
		j++
	}
	return result
}

// func main() {
// 	lst := []int{4, 2, 5, 3, 1, 7, 6}
// 	fmt.Println(mergeSort(lst))
// }
