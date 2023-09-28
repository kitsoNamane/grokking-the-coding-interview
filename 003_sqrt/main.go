package main

import "fmt"

func mySqrtNaive(x int) int {
	if x < 2 {
		return x
	}

	sqrt := 0

	for i := 1; i < x+1; i++ {
		_sqrt := i * i

		if _sqrt > x {
			return sqrt
		}

		sqrt = i
	}

	return 0
}

func mySqrtBinarySearch(x int) int {
	if x < 2 {
		return x
	}

	left, right := 2, x/2
	pivot := 0
	var num float64

	for left <= right {
		pivot = left + (right-left)/2
		num = float64(pivot) * float64(pivot)
		if num > float64(x) {
			right = pivot - 1
		} else if num < float64(x) {
			left = pivot + 1
		} else {
			return pivot
		}
	}

	return right
}

func main() {
	fmt.Println(2 == mySqrtNaive(8))
	fmt.Println(1 == mySqrtNaive(1))
	fmt.Println(2 == mySqrtNaive(4))
	fmt.Println(4 == mySqrtNaive(16))
	fmt.Println(4 == mySqrtNaive(19))
}
