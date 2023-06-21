package iohtml

import (
	"fmt"
	"io/fs"
	"os"

	"golang.org/x/net/html"
)

// GetHTMLDoc открывает html-файл и возвращает дерево html-документа.
func GetHTMLDoc(filename string) (doc *html.Node, err error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, fs.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("getHTMLDoc: %q", err)
	}
	defer f.Close()
	return html.Parse(f)
}

// ForEachNode вызывает функции pre(x) и post(x) для каждого узла х
// в дереве с корнем n. Обе функции необязательны.
// рге вызывается до посещения дочерних узлов, a post - после.
func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
