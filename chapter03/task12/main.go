package main

import (
	"fmt"
)

func main() {
	s1 := "12345🤣6789🤣"
	s2 := "🤣6789🤣12345"
	fmt.Printf("%s is anagram %s. It's %v\n", s1, s2, isAnagram(s1, s2))
	s1 = "12345🤣6789"
	s2 = "12345678😂9"
	fmt.Printf("%s is anagram %s. It's %v\n", s1, s2, isAnagram(s1, s2))
}

func isAnagram(s1 string, s2 string) bool {
	r1 := []rune(s1)
	r2 := []rune(s2)
	lenr := len(r1)
	if lenr != len(r2) {
		return false
	}
	for i, v := range r1 {
		if !containsRune(r2[:lenr-i], v) {
			return false
		}
	}
	return true
}

func containsRune(ar []rune, r rune) bool {
	for i, v := range ar {
		if r == v {
			if ar[i] != ar[len(ar)-1] {
				ar[i] = ar[len(ar)-1]
			}
			return true
		}

	}
	return false
}
