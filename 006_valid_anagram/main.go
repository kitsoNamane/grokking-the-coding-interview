package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	Items   []rune
	topIdex int
}

func NewStack(items []rune) *Stack {
	return &Stack{Items: items, topIdex: len(items) - 1}
}

func (s *Stack) Pop() rune {
	if s.topIdex < 0 {
		return 0
	}

	if s.topIdex >= 0 {
		popedItem := s.Items[s.topIdex]
		s.Items = s.Items[0:s.topIdex]
		s.topIdex = s.topIdex - 1
		return popedItem
	}

	return 0
}

func isAnagramStack(s, t string) bool {
	_s := []rune(s)
	sort.Slice(_s, func(i, j int) bool {
		return _s[i] < _s[j]
	})
	_sStack := NewStack(_s)

	_t := []rune(t)
	sort.Slice(_t, func(i, j int) bool {
		return _t[i] < _t[j]
	})
	_tStack := NewStack(_t)

	if len(_t) < len(_s) {
		return false
	}

	for range _t {
		if _tStack.Pop() != _sStack.Pop() {
			return false
		}
	}

	return true
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make(map[rune]int64, len(s))

	for i := range s {
		counter[rune(s[i])]++
		counter[rune(t[i])]--
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("hello", "world"))
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("aacc", "ccac"))
	fmt.Println(isAnagram("a", "b"))
	fmt.Println(isAnagram("aa", "a"))
}
