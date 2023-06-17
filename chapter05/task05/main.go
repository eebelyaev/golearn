package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://kosmodrive.com/ru/rostov-na-donu/"
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "main: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Words: %d, Images: %d", words, images)
}

// CountWordsAndlmages выполняет HTTP-запрос GET HTML-документа
// url и возвращает количество слов и изображений в нем.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				images++
				break
			}
		}
	} else if n.Type == html.TextNode {
		scanner := bufio.NewScanner(strings.NewReader(n.Data))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words++
		}
	}
	w, i := countWordsAndImages(n.FirstChild)
	images += i
	words += w
	w, i = countWordsAndImages(n.NextSibling)
	images += i
	words += w
	return
}
