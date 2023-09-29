package main

import (
	"fmt"
	"strings"
)

func reverseVowelsNaive(s string) string {
	vowels := "aeiouAEIOU"

	var reversed []string
	var vowelStack []string

	for _, letter := range s {
		if strings.ContainsRune(vowels, letter) {
			reversed = append(reversed, "")
			vowelStack = append(vowelStack, string(letter))
			continue
		}
		reversed = append(reversed, string(letter))
	}

	popIndex := len(vowelStack) - 1
	for i, letter := range reversed {
		if letter == "" {
			reversed[i] = vowelStack[popIndex]
			popIndex--
		}
	}

	return strings.Join(reversed, "")
}

func reverseVowelsTwoPointer(s string) string {
	vowels := "aeiouAEIOU"
	first, last := 0, len(s)-1

	array := []rune(s)
	for first < last {
		for first < last && !strings.ContainsRune(vowels, array[first]) {
			first++
		}

		for first < last && !strings.ContainsRune(vowels, array[last]) {
			last--
		}

		temp := array[first]
		array[first] = array[last]
		array[last] = temp
		first++
		last--
	}

	return string(array)
}

func main() {

	fmt.Println("holle" == reverseVowelsTwoPointer("hello"))
	fmt.Println("AEIOU" == reverseVowelsTwoPointer("UOIEA"))
	fmt.Println("DusUgnGires" == reverseVowelsTwoPointer("DesignGUrus"))
}
