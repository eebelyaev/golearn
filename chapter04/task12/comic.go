package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Comic struct {
	Num        int
	Img        string
	Title      string
	Transcript string
}

func (c *Comic) print() {
	t := c.Transcript
	posB := strings.Index(t, "[[") + len("[[")
	posE := strings.Index(t[posB:], "]]") + posB
	if posB >= 0 && posE > posB {
		t = t[posB:posE]
	}
	fmt.Printf("%d\tURL: \"%s\"\n\t\"%s\"\n", c.Num, c.Img, t)
}

func loadComic(n int) (*Comic, error) {
	resp, err := http.Get(fmt.Sprintf(ComicURLTempl, n))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var c Comic
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		log.Fatalf("Сбой демаршалинга JSON #%d: %s", n, err)
		return nil, err
	}
	return &c, nil
}
