package main

import "fmt"

func pangram(sentence string) bool {

	set := make(map[rune]bool)
	counter := 0

	for _, v := range sentence {
		if v == 32 {
			continue
		}

		if v > 90 {
			v = v - 32
		}

		if v < 65 || v > 122 {
			continue
		}

		if _, ok := set[v]; ok == false {
			counter++
		}
		set[v] = true
	}

	if counter == 26 {
		return true
	}

	return false
}

func main() {

	sentence := "thequickbrownfoxjumpsoverthelazydog"

	fmt.Println(pangram(sentence))

}
