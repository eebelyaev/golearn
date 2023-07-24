package main

import (
	"errors"
	"fmt"
	"golearn/chapter05/links"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const SAVED_DOMAIN = "https://go.dev/"

func main() {
	// Поиск в ширину, начиная с аргумента командной строки.
	breadthFirst(crawl, []string{"https://golang.org"})
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	} else {
		for _, u := range list {
			if strings.HasPrefix(u, SAVED_DOMAIN) {
				savePage(u)
			}
		}
	}
	return list
}

// savePage сохраняет указанную страницу в соответствующем каталоге.
// Если каталога не существует, он создается.
func savePage(url string) error {
	path, _ := strings.CutPrefix(url, SAVED_DOMAIN)
	if path = strings.Trim(path, "/"); len(path) > 0 {
		path = "res/" + path
	} else {
		path = "res"
	}
	// проверим отсутствие файла
	if _, err := os.Stat(path + "/index.html"); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(path, 0750)
		if err != nil && !os.IsExist(err) {
			return err
		}

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("получение %s: %s", url, resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return os.WriteFile(path+"/index.html", body, 0644)
	}
	return nil
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
