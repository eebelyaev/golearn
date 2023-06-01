package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	sb := []byte("☀️ Hello! 😎 Улыбок тебе дед Макар! 🤣")
	print(sb)
	reverse(sb)
	print(sb)
}

// reverse обращает порядок чисел "на месте"
func reverse(sb []byte) {
	for i, j := 0, len(sb); i < j; {
		sbl := sb[i:j]
		ri, si := utf8.DecodeRune(sbl)
		rj, sj := utf8.DecodeLastRune(sbl)
		if ri != rj {
			var dst, src []byte
			if si < sj {
				dst, src = sbl[sj:], sbl[si:len(sbl)-(sj-si)]
			} else if si > sj {
				dst, src = sbl[sj:len(sbl)-(si-sj)], sbl[si:]
			}
			copy(dst, src)
			copy(sbl[:sj], []byte(string(rj)))
			copy(sbl[len(sbl)-si:], []byte(string(ri)))
		}
		i += sj
		j -= si
	}
}

func print(sb []byte) {
	fmt.Printf("%q\n", sb)
	//fmt.Printf("%v\n", sb)
}
