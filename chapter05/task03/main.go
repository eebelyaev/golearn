package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

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
	for _, t := range visitText(nil, doc) {
		fmt.Println(t)
	}
}

// visitMap добавляет в m все ссылки,
// найденные в n, и возвращает результат.
func visitText(list []string, n *html.Node) []string {
	if n == nil {
		return list
	}
	if n.Type == html.TextNode && strings.TrimSpace(n.Data) != "" {
		list = append(list, n.Data)
	}
	if n.Type != html.ElementNode || n.Data != "script" && n.Data != "style" {
		list = visitText(list, n.FirstChild)
		list = visitText(list, n.NextSibling)
	}
	return list
}

func loadHTML(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("loadHTML: %q", err)
	}
	return file, nil
}
