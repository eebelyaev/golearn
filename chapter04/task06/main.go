package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

const asciiSpace = ' '

func main() {
	sb := []byte("🤣\nHello\t\t🤣")
	print(sb)
	sb = removeSpaces(sb)
	print(sb)
}

func removeSpaces(sb []byte) []byte {
	firstSpace := -1
	for i := 0; i < len(sb); {
		r, size := utf8.DecodeRuneInString(string(sb[i:]))
		if unicode.IsSpace(r) {
			if firstSpace < 0 {
				firstSpace = i
			}
		} else {
			if firstSpace >= 0 {
				sb[firstSpace] = asciiSpace
				sb = append(sb[:firstSpace+1], sb[i:]...)
				firstSpace = -1
			}
		}
		i += size
	}
	if firstSpace >= 0 {
		sb = append(sb[:firstSpace], asciiSpace)
	}
	return sb
}

func print(sb []byte) {
	fmt.Printf("%q\n", sb)
	fmt.Printf("%v\n", sb)
}
