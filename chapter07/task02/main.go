package main

import (
	"fmt"
	"io"
	"os"
)

type cwriter struct {
	w io.Writer
	n int64
}

func (cw *cwriter) Write(p []byte) (n int, err error) {
	n, err = cw.w.Write(p)
	cw.n += int64(n)
	return
}

func main() {
	wn, n := CountingWriter(os.Stdout)
	fmt.Fprintln(wn, "1234567")
	fmt.Fprintln(wn, *n)
	fmt.Println(*n)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := cwriter{w, 0}
	return &cw, &cw.n
}
