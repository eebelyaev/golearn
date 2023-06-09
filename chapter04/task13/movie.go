package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const MovieServiceURL = "http://www.omdbapi.com/?i=tt3896198&apikey=6dbbd51d"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func (m *Movie) print() {
	fmt.Println(m)
}

func loadMovie(t string) (*Movie, error) {
	s := MovieServiceURL + "&t=" + url.QueryEscape(t)
	resp, err := http.Get(s)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сбой запроса: %s", resp.Status)
	}
	var m Movie
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Movie) savePoster() error {
	if m == nil || m.Poster == "" {
		return fmt.Errorf("savePoster:no poster for movie")
	}
	resp, err := http.Get(m.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(m.Title + ".jpg")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("Success!")
	return nil
}
