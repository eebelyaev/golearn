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
	printMap(visitMap(nil, doc))
}

// visitMap добавляет в m все ссылки,
// найденные в n, и возвращает результат.
func visitMap(m map[string][]string, n *html.Node) map[string][]string {
	if m == nil {
		m = make(map[string][]string, 0)
	}
	if n == nil {
		return m
	}
	if n.Type == html.ElementNode {
		if list, ok := appendAttr(m[n.Data], n); ok {
			m[n.Data] = list
		}
	}
	m = visitMap(m, n.FirstChild)
	m = visitMap(m, n.NextSibling)
	return m
}

func appendAttr(list []string, n *html.Node) ([]string, bool) {
	var attrName string
	switch n.Data {
	case "a", "link":
		attrName = "href"
	case "img", "script":
		attrName = "src"
	default:
		return nil, false
	}
	for _, a := range n.Attr {
		if a.Key == attrName {
			return append(list, a.Val), true
		}
	}
	return nil, false
}

func loadHTML(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("loadHTML: %q", err)
	}
	return file, nil
}

func printMap(m map[string][]string) {
	for k, v := range m {
		fmt.Println(k)
		for _, t := range v {
			fmt.Println(t)
		}
		fmt.Println()
	}
}
