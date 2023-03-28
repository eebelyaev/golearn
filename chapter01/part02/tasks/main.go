package main

import (
	"fmt"
	"os"
	"time"
	// "strings"
)

func main() {
	task01()
	task02()
	task03()
}

func task01() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func task02() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}

func task03() {
	fmt.Println(time.Now())
}
