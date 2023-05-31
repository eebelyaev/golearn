package main

import (
	"fmt"
)

func main() {
	s := []string{"1", "2", "3", "3", "4", "5", "5", "5", "5", "5", "5", "5", "1"}
	fmt.Printf("%v, %T\n", s, s)
	for i, j := 0, 0; i < len(s)-1; i++ {
		for j = i + 1; j < len(s) && s[i] == s[j]; j++ {
		}
		if j--; i != j {
			s = append(s[:i], s[j:]...)
		}
	}
	fmt.Printf("%v, %T\n", s, s)
}
