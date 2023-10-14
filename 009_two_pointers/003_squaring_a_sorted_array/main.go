package main

import (
	"fmt"
)

func makeSquares(arr []int) []int {
	n := len(arr)
	squares := make([]int, n, n)
	highestSquareIdx := len(squares)
	left := 0
	right := len(squares) - 1

	for left <= right {
		leftSquare := arr[left] * arr[left]
		rightSquare := arr[right] * arr[right]

		if highestSquareIdx > 0 {
			highestSquareIdx--
		}

		if leftSquare > rightSquare {
			squares[highestSquareIdx] = leftSquare
			left++
		} else {
			squares[highestSquareIdx] = rightSquare
			right--
		}
	}

	return squares
}

func main() {
	fmt.Println(makeSquares([]int{-2, -1, 0, 2, 3}))
	fmt.Println(makeSquares([]int{-3, -1, 0, 1, 2}))
	fmt.Println(makeSquares([]int{-4, -2, 0, 2, 4}))
	fmt.Println(makeSquares([]int{-5, -4, -3, -2, -1}))
	fmt.Println(makeSquares([]int{-1, 0, 1, 2, 3}))
}
