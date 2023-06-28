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
	id := "fid"
	found := ElementByID(doc, id)
	fmt.Printf("id = %s, element = %v", id, found)
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return true
				}
			}
		}
		return false
	}
	return forEachNode(doc, pre)
}

// forEachNode вызывает функции pre(x) и post(x) для каждого узла х
// в дереве с корнем n.
// рге вызывается до посещения дочерних узлов.
func forEachNode(n *html.Node, pre func(n *html.Node) bool) *html.Node {
	if pre(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nn := forEachNode(c, pre)
		if nn != nil {
			return nn
		}
	}
	return nil
}
