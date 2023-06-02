package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	input := "☀️ Hello! 😎 Улыбок тебе дед Макар! 🤣 ДЕД"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	wordFreq := make(map[string]int)
	for scanner.Scan() {
		wordFreq[strings.ToLower(scanner.Text())]++
	}
	print(wordFreq)
}

func print(m map[string]int) {
	fmt.Print("\nword\tcount\n")
	for k, v := range m {
		if v > 1 {
			fmt.Printf("%s\t%d\n", k, v)
		}
	}
}
