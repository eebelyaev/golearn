package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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
		sep = "\n"
	}
	fmt.Println(s)
}

func task02() {
	for i, arg := range os.Args {
		fmt.Printf("%d\t%s\n", i, arg)
	}
}

func task03() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		task01()
	}
	durationOfTask01 := time.Since(start).Milliseconds()

	start = time.Now()
	for i := 0; i < 10000; i++ {
		task02()
	}
	durationOfTask02 := time.Since(start).Milliseconds()

	start = time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Println(strings.Join(os.Args, ", "))
	}
	durationOfTask03 := time.Since(start).Milliseconds()

	fmt.Printf("Duration of task1 execution: %d ms\n", durationOfTask01)
	fmt.Printf("Duration of task2 execution: %d ms\n", durationOfTask02)
	fmt.Printf("Duration of task3 execution: %d ms", durationOfTask03)
}
