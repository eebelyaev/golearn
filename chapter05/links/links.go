// Пакет links предоставляет функцию для извлечения ссылок.
package links

import (
	"fmt"
	"golearn/chapter05/iohtml"

	"golang.org/x/net/html"

	"net/http"
)

// Extract выполняет HTTP-запрос GET по определенному URL, выполняет
// синтаксический анализ HTML и возвращает ссылки в HTML-документе.
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("получение %s: %s", url, resp.Status)
	}
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
