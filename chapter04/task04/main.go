package main

import (
	"fmt"
)

func main() {
	a := [...]int{1: 1, 2, 3, 4, 5}
	fmt.Printf("%v, %T\n", a, a)
	rotate(a[:], 5)
	fmt.Printf("%v, %T\n", a, a)
}

// Циклический сдвиг влево на sh позиций
func rotate(a []int, sh int) {
	sh = sh % len(a)
	if sh < 0 {
		sh = len(a) + sh
	}
	if sh > 0 {
		s := make([]int, sh)
		copy(s, a[:sh])
		copy(a[0:], a[sh:])
		copy(a[len(a)-sh:], s)
	}
}
