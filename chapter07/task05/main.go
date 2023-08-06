package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := "0123456789"
	sz := int64(9)
	r0 := io.LimitReader(strings.NewReader(s), sz)
	r := LimitReader(strings.NewReader(s), sz)
	print(r0, s)
	print(r, s)
}

func print(r io.Reader, s string) {
	fmt.Printf("%T =====================\n", r)
	buf := make([]byte, 5)
	for n, err := r.Read(buf); err != io.EOF; n, err = r.Read(buf) {
		if err != nil {
			fmt.Printf("ошибка чтения: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", buf[:n])
	}
}
