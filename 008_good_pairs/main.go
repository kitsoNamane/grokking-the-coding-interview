package main

import "fmt"

func numGoodPairs(nums []int) int {
	pairCount := 0

	for i := range nums {
		for j := range nums {
			if nums[i] == nums[j] && i < j {
				pairCount++
			}
		}
	}

	return pairCount
}

func main() {

	fmt.Println(4 == numGoodPairs([]int{1, 2, 3, 1, 1, 3}))
	fmt.Println(6 == numGoodPairs([]int{1, 1, 1, 1}))
	fmt.Println(0 == numGoodPairs([]int{1, 2, 3}))
}
