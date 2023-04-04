package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	urlNames := []string{"http://ya.ru", "http://google.com", "http://google.ru"}
	var urlFound string
	var sizeFound int64
	for _, url := range urlNames {
		if a := getBodySize(url); a > sizeFound {
			urlFound, sizeFound = url, a
		}
	}
	fmt.Printf("%v\t%d\n", urlFound, sizeFound)
}

func getBodySize(url string) int64 {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	w, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("Duration of get %s: %d\n", url, time.Since(start).Milliseconds())
	return w
}
