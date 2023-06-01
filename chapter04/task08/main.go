package main

import (
	"fmt"
	"unicode"
)

const (
	catControl = "Control"
	catDigit   = "Digit"
	catGraphic = "Graphic"
	catLetter  = "Letter"
	catLower   = "Lower"
	catMark    = "Mark"
	catNumber  = "Number"
	catPunct   = "Punct"
	catSpace   = "Space"
	catSymbol  = "Symbol"
	catTitle   = "Title"
	catUpper   = "Upper"
)

var fncCats map[string]func(r rune) bool

func init() {
	fncCats = map[string]func(r rune) bool{
		catControl: unicode.IsControl,
		catDigit:   unicode.IsDigit,
		catGraphic: unicode.IsGraphic,
		catLetter:  unicode.IsLetter,
		catLower:   unicode.IsLower,
		catMark:    unicode.IsMark,
		catNumber:  unicode.IsNumber,
		catPunct:   unicode.IsPunct,
		catSpace:   unicode.IsSpace,
		catSymbol:  unicode.IsSymbol,
		catTitle:   unicode.IsTitle,
		catUpper:   unicode.IsUpper,
	}
}

func main() {
	s := "☀️ Hello! 😎 Улыбок тебе дед Макар! 🤣"
	fmt.Println(s)
	printStat(s)
}

func printStat(s string) {
	cntCats := make(map[string]int)
	invalid := 0 // Количество некорректных символов UTF-8
	for _, r := range s {
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		for k, f := range fncCats {
			if f(r) {
				cntCats[k]++
			}
		}
	}
	fmt.Print("\ncategory\tcount\n")
	for k, v := range cntCats {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d неверных символов UTF-8\n", invalid)
	}
}
