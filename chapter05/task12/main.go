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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	newOutline(doc)
}

func newOutline(n *html.Node) {
	var stack []string
	pre := func(n *html.Node) {
		if n.Type == html.ElementNode {
			stack = append(stack, n.Data) // Внесение дескриптора в стек
			fmt.Println(stack)
		}
	}
	post := func(n *html.Node) {
		if n.Type == html.ElementNode {
			stack = stack[:len(stack)-1]
		}
	}
	forEachNode(n, pre, post)
}

// forEachNode вызывает функции pre(x) и post(x) для каждого узла x
// в дереве с корнем п. Обе функции необязательны.
// рге вызывается до посещения дочерних узлов, a post - после,
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
