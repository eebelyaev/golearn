package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	f, n, err := fetch("http://ya.ru")
	fmt.Printf("f = %s, n = %d, err = %q", f, n, err)
}

// Fetch загружает URL и возвращает имя и длину локального файла.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// Закрытие файла; если есть ошибка Сору, возвращаем ее.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
