package main

import (
	"fmt"
	"math"
)

func shortestDistance(words []string, wordOne, wordTwo string) int {
	var shortestDistance = len(words)
	var firstPivotIndex = -1
	var secondPivotIndex = -1

	for i := range words {
		if words[i] == wordOne {
			firstPivotIndex = i
		}

		if words[i] == wordTwo {
			secondPivotIndex = i
		}

		_shortestDistance := int(math.Abs(float64(secondPivotIndex - firstPivotIndex)))
		if firstPivotIndex >= 0 && secondPivotIndex >= 0 && _shortestDistance <= shortestDistance {
			shortestDistance = _shortestDistance
		}
	}

	return shortestDistance
}

func main() {
	fmt.Println(5 == shortestDistance([]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}, "fox", "dog"))

	fmt.Println(1 == shortestDistance([]string{"a", "c", "d", "b", "a"}, "a", "b"))

	fmt.Println(4 == shortestDistance([]string{"a", "b", "c", "d", "e"}, "a", "e"))

	fmt.Println(1 == shortestDistance([]string{"repeated", "words", "in", "the", "array", "words"}, "repeated", "words"))

}
