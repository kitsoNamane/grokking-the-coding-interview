package main

import (
	"fmt"
	"sort"
)

func containsDuplicates_BruteForce(nums []int) bool {
	nums_len := len(nums)
	for i := 0; i < nums_len; i++ {
		for j := i + 1; j < nums_len; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func containsDuplicates_Set(nums []int) bool {
	unique_set := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := unique_set[v]; ok {
			return true
		}

		unique_set[v] = true
	}

	return false
}

func containsDuplicates_Sorting(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

func main() {
	numsOne := []int{1, 2, 3, 1}
	numsTwo := []int{1, 2, 3, 4}
	numsThree := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicates_Sorting(numsOne))
	fmt.Println(containsDuplicates_Sorting(numsTwo))
	fmt.Println(containsDuplicates_Sorting(numsThree))
}
