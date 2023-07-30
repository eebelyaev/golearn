package main

import (
	"fmt"
)

func main() {
	x := 5
	fmt.Printf("plusOne(%d) = %d\n", x, plusOne(x))
}

func plusOne(x int) (res int) {
	defer func() {
		p := recover()
		res = p.(int)
	}()
	func(x int) {
		panic(x + 1)
	}(x)
	return
}
