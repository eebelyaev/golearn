package main

import (
	"fmt"
	"os"
	"strings"

	"golearn/chapter05/iohtml"

	"golang.org/x/net/html"
)

var depth = 0

func main() {
	doc, err := iohtml.GetHTMLDoc("..\\task04\\example.html")
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	iohtml.ForEachNode(doc, startElement, endElement)
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	} else if n.Type == html.TextNode || n.Type == html.CommentNode {
		if s := strings.TrimSpace(n.Data); s != "" {
			fmt.Printf("%*s%s\n", depth*2, "", s)
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	} else if n.Type == html.CommentNode || n.Type == html.TextNode {
		depth--
	}
}
