package main

import (
	"fmt"
	"io/fs"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := loadHTML("example.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visitMap(nil, doc) {
		fmt.Println(k, ": ", v)
	}
}

// visitMap добавляет в m все ссылки,
// найденные в n, и возвращает результат.
func visitMap(m map[string]int, n *html.Node) map[string]int {
	if m == nil {
		m = make(map[string]int, 0)
	}
	if n == nil {
		return m
	}
	if n.Type == html.ElementNode {
		m[n.Data]++
	}
	m = visitMap(m, n.FirstChild)
	return visitMap(m, n.NextSibling)
}

func loadHTML(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("loadHTML: %q", err)
	}
	return file, nil
}
