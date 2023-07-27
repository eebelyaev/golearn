package main

import (
	"fmt"
	"golearn/chapter05/iohtml"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := iohtml.GetHTMLDoc("..\\task04\\example.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	images := ElementsByTagName(doc, "img")
	fmt.Println(len(images))
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println(len(headings))
}

func ElementsByTagName(doc *html.Node, names ...string) (res []*html.Node) {
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, name := range names {
				if n.Data == name {
					res = append(res, n)
				}
			}
		}
	}
	iohtml.ForEachNode(doc, pre, nil)
	return
}
