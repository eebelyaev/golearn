package main

import (
	"fmt"
	"golearn/chapter05/iohtml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

const SAVED_DOMAIN = "https://go.dev/"

func main() {
	// Поиск в ширину, начиная с аргумента командной строки.
	breadthFirst(crawl, []string{"https://golang.org"})
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func savePage(resp *http.Response, url string) error {
	if strings.HasPrefix(url, SAVED_DOMAIN) {
		path, _ := strings.CutPrefix(url, SAVED_DOMAIN)
		path = "res/" + strings.Trim(path, "/")
		err := os.MkdirAll(path, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(path+"/index.html", body, 0644)
	}
	return nil
}

// extract выполняет HTTP-запрос GET по определенному URL, выполняет
// синтаксический анализ HTML и возвращает ссылки в HTML-документе.
func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("получение %s: %s", url, resp.Status)
	}
	savePage(resp, url)
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("анализ %s как HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // Игнорируем некорректные URL
				}
				links = append(links, link.String())
			}
		}
	}
	iohtml.ForEachNode(doc, visitNode, nil)
	return links, nil
}

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
