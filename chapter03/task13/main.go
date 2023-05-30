package main

import (
	"fmt"
)

const (
	x, KB, MB, GB, TB, PB, EB, ZB, YB = 1000, x, KB * x, MB * x, GB * x, TB * GB, PB * x, EB * x, ZB * x
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(YB / 1000000000000)
}
