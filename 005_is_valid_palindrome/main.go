package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	sLower := []rune(strings.ToLower(s))
	first, last := 0, len(sLower)-1

	for first <= last {
		left := sLower[first]
		right := sLower[last]

		isLetter := left >= 97 && left <= 122
		isNumber := left >= 48 && left <= 57

		if isLetter && isNumber && left != right {
			return false
		}

		first++
		last--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal, Panama!"))
}
