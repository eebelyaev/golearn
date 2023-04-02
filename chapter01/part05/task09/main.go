package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func main() {
	urlNames := []string{"http://ya.ru", "http://google.com", "asot.ru"}
	for _, url := range urlNames {
		fmt.Println(url)
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\nСтатус %s: %v\n", url, resp.Status)
	}
}
