package syan

import (
	"fmt"
	"golearn/chapter05/iohtml"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type LinkReader struct {
	links []string
}

func (lr *LinkReader) Read(p []byte) (n int, err error) {
	if len(lr.links) > 0 {
		v := lr.links[0]
		if len(p) >= len(v) {
			n = copy(p, v)
			lr.links = lr.links[1:]
		} else {
			err = fmt.Errorf("buffer too small: need %d bytes", len(v))
		}
	} else {
		err = io.EOF
	}
	return
}

func NewReader(s string) (r io.Reader) {
	var lr LinkReader
	lr.links, _ = ExtractFromString(s)
	return &lr
}

// ExtractFromString выполняет синтаксический анализ HTML и
// возвращает ссылки в HTML-документе.
func ExtractFromString(body string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("анализ как HTML: %v", err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				links = append(links, a.Val)
			}
		}
	}
	iohtml.ForEachNode(doc, visitNode, nil)
	return links, nil
}
