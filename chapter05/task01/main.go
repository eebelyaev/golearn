package main

import (
	"fmt"
	"io/fs"
	"os"

	"golang.org/x/net/html"
)

func main() {
	//doc, err := html.Parse(os.Stdin)
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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit добавляет в links все ссылки,
// найденные в n, и возвращает результат.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}

func loadHTML(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("loadHTML: %q", err)
	}
	return file, nil
}
