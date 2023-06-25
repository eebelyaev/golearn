package main

import (
	"os"
	"testing"
	"time"

	"golang.org/x/net/html"
)

var tests = []struct {
	input                *html.Node
	depthStart, depthEnd int
	wantStart, wantEnd   string
}{
	{
		&html.Node{
			Type: html.ElementNode,
			Data: "img",
			Attr: []html.Attribute{
				{Key: "src", Val: "ex.png"},
				{Key: "width", Val: "100"},
			},
		},
		2,
		0,
		"  <img src=\"ex.png\" width=\"100\"/>\n",
		"",
	},
	{
		&html.Node{
			Type: html.ElementNode,
			Data: "img",
			Attr: []html.Attribute{
				{Key: "src", Val: "ex.png"},
				{Key: "width", Val: "100"},
			},
			FirstChild: &html.Node{
				Type: html.TextNode,
				Data: "Hello",
			},
		},
		2,
		0,
		"  <img src=\"ex.png\" width=\"100\">\n",
		"</img>\n",
	},
	{
		&html.Node{
			Type: html.DocumentNode,
			Data: "",
			FirstChild: &html.Node{
				Type: html.ElementNode,
				Data: "img",
			},
		},
		1,
		1,
		"",
		"",
	},
}

func TestStartElement(t *testing.T) {
	for i, test := range tests {
		depth = 1
		time.Sleep(time.Second)
		got, _ := captureOut(startElement, test.input)
		if got != test.wantStart {
			t.Errorf("startElement(%d) = %q, требуется %q", i, got, test.wantStart)
		} else if depth != test.depthStart {
			t.Errorf("endElement(%d) depth = %d, требуется %d", i, depth, test.depthStart)
		}
	}
}
func TestEndElement(t *testing.T) {
	for i, test := range tests {
		depth = 1
		time.Sleep(time.Second)
		got, _ := captureOut(endElement, test.input)
		if got != test.wantEnd {
			t.Errorf("endElement(%d) = %q, требуется %q", i, got, test.wantEnd)
		} else if depth != test.depthEnd {
			t.Errorf("endElement(%d) depth = %d, требуется %d", i, depth, test.depthEnd)
		}
	}
}

func TestStartElement10(t *testing.T) {
	for i := 0; i < 20; i++ {
		TestStartElement(t)
	}
}

func TestEndElement10(t *testing.T) {
	for i := 0; i < 20; i++ {
		TestEndElement(t)
	}
}

func captureOut(f func(*html.Node), n *html.Node) (string, error) {
	defer func(stdout *os.File) {
		os.Stdout = stdout
	}(os.Stdout)
	out, err := os.CreateTemp("", "stdout")
	if err != nil {
		return "", err
	}
	defer out.Close()
	outname := out.Name()
	os.Stdout = out

	f(n)

	err = out.Close()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(outname)
	if err != nil {
		return "", err
	}
	os.Remove(outname)
	return string(data), nil
}
