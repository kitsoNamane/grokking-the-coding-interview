package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	sLower := []rune(strings.ToLower(s))
	first, last := 0, len(sLower)-1

	for first < last {

		for first < last && !isAlphaNumeric(sLower[first]) {
			first++
		}

		for first < last && !isAlphaNumeric(sLower[last]) {
			last--
		}

		left := sLower[first]
		right := sLower[last]

		if left != right {
			fmt.Printf("left %v - index %v\n", left, first)
			fmt.Printf("right %v - index %v\n", right, last)
			return false
		}

		first++
		last--
	}

	return true
}

func isAlphaNumeric(r rune) bool {
	isLetter := r >= 97 && r <= 122
	isNumber := r >= 48 && r <= 57

	return isLetter || isNumber
}

func main() {
	fmt.Println(isPalindrome("ra aar"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome(" "))
}
