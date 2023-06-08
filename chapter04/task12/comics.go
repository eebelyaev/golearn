package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
)

const (
	ComicURLTempl = "https://xkcd.com/%d/info.0.json"
	ComicCount    = 2784
	CashFile      = "xkcd.json"
)

type Comics []*Comic

func newComics() Comics {
	return make(Comics, 0, ComicCount)
}

func (comics Comics) load(n int) Comics {
	comics = comics.loadFromCash()
	lenComics := len(comics)
	if lenComics >= ComicCount {
		return comics
	}

	if n > ComicCount-lenComics {
		n = ComicCount - lenComics
	}
	for i := lenComics; i < lenComics+n && i < ComicCount; i++ {
		if c, err := loadComic(i + 1); err != nil {
			log.Printf("%d comics are loaded\n", i)
			return comics
		} else {
			comics = append(comics, c)
		}
	}
	log.Printf("%d comics are loaded\n", n)
	return comics
}

func (comics Comics) loadFromCash() Comics {
	data, err := os.ReadFile(CashFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "loadFromCash: %v\n", err)
		return comics
	}
	if err := json.Unmarshal(data, &comics); err != nil {
		log.Fatalf("Сбой демаршалинга JSON: %s", err)
	}
	return comics
}

func (comics Comics) save() {
	if len(comics) == 0 {
		return
	}
	if data, err := json.Marshal(comics); err != nil {
		log.Fatalf("Сбой маршалинга JSON: %s", err)
	} else {
		os.WriteFile(CashFile, data, fs.ModePerm)
	}
}
