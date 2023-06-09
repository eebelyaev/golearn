package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	if m, err := loadMovie(strings.Join(os.Args[1:], " ")); err != nil {
		log.Fatalf("main:%s", err)
	} else {
		m.print()
		m.savePoster()
	}
}
