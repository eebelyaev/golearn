package main

import (
	"fmt"
	"io"
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
		w, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s %v %d\n", url, err, w)
			os.Exit(1)
		}
		fmt.Printf("\nДля сайта %s скопировано %d байт", url, w)
	}
}
