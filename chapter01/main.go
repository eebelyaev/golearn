package main

import "fmt"

func main() {
	a := 25
	var b = uint8(1 + int8(5))
	fmt.Printf("%d %d", a, b)
}
