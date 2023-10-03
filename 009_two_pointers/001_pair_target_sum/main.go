package main

import (
	"fmt"
	"slices"
)

func searchTwoPair(arr []int, targetSum int) []int {
	firstIndex := 0
	secondIndex := len(arr) - 1

	for firstIndex < secondIndex {

		sum := arr[firstIndex] + arr[secondIndex]

		if sum < targetSum {
			firstIndex++
		} else if sum > targetSum {
			secondIndex--
		} else if sum == targetSum {
			return []int{firstIndex, secondIndex}
		}
	}

	return []int{-1, -1}
}

func search(arr []int, targetSum int) []int {
	nums := make(map[int]int)

	for i := range arr {
		if yIndex, ok := nums[targetSum-arr[i]]; ok {
			return []int{i, yIndex}
		} else {
			nums[arr[i]] = i
		}
	}

	return []int{-1, -1}
}

func check(expected, actual []int) bool {

	for _, x := range expected {
		if !slices.Contains(actual, x) {
			return false
		}
	}
	return true
}

func main() {
	indexPairsOne := []int{1, 3}
	fmt.Println(check(indexPairsOne, search([]int{1, 2, 3, 4, 6}, 6)))
	indexPairsTwo := []int{0, 2}
	fmt.Println(check(indexPairsTwo, search([]int{2, 5, 9, 11}, 11)))
}
