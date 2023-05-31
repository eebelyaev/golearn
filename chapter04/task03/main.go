package main

import (
	"fmt"
	"golearn/chapter04/rev"
)

func main() {
	print()
}

func print() {
	a := [...]int{1: 1, 2, 3, 5: 5}
	fmt.Printf("%v, %T\n", a, a)
	rev.Reverse(a[:])
	fmt.Printf("%v, %T\n", a, a)
	a = [...]int{1: 1, 2, 3, 5: 5}
	Reverse(&a)
	fmt.Printf("%v, %T\n", a, a)
}

// reverse обращает порядок чисел "на месте"
func Reverse(p *[6]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}
