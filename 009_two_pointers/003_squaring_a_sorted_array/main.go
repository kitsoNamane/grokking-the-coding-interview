package main

import (
	"fmt"
	"slices"
)

func makeSquares(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] < 0 {
			arr[i] = -1 * arr[i]
		}
		arr[i] = arr[i] * arr[i]
	}

	slices.Sort(arr)

	return arr
}

func main() {
	fmt.Println(makeSquares([]int{-2, -1, 0, 2, 3}))
	fmt.Println(makeSquares([]int{-3, -1, 0, 1, 2}))
	fmt.Println(makeSquares([]int{-4, -2, 0, 2, 4}))
	fmt.Println(makeSquares([]int{-5, -4, -3, -2, -1}))
	fmt.Println(makeSquares([]int{-1, 0, 1, 2, 3}))
}
