package main

import "fmt"

func remove(arr []int) int {
	first, second := 0, 1

	for first < len(arr)-2 {
		for arr[first] == arr[second] && second != len(arr)-1 {
			second++
		}

		temp := arr[first+1]
		arr[first+1] = arr[second]
		arr[second] = temp

		if arr[first] != arr[second] {
			first++
			second = first + 1
		}

		fmt.Printf("first: %d - second: %d\n", first, second)

		fmt.Println(arr)
	}

	return 0
}

func main() {
	fmt.Println(2 == remove([]int{2, 2, 2, 11}))
	fmt.Println(4 == remove([]int{2, 3, 4, 4, 6, 9, 9}))
}
