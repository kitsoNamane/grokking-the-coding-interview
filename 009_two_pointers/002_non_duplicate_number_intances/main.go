package main

import "fmt"

func remove(arr []int) int {
	nextNonDuplicateIndex := 1

	for i := 0; i < len(arr); i++ {
		if arr[nextNonDuplicateIndex-1] != arr[i] {
			arr[nextNonDuplicateIndex] = arr[i]
			nextNonDuplicateIndex++
		}
	}

	return nextNonDuplicateIndex
}

func main() {
	fmt.Println(2 == remove([]int{2, 2, 2, 11}))
	fmt.Println(4 == remove([]int{2, 3, 3, 3, 6, 9, 9}))
	fmt.Println(4 == remove([]int{1, 1, 2, 2, 3, 3, 4, 4}))
}
