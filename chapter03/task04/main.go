/*
Упражнение 3.4. Создайте веб-сервер, который вычисляет поверхности и
возвращает клиенту SVG-данные.
*/
package main

import (
	"golearn/chapter03/task04/draw"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	for k, v := range r.URL.Query() {
		switch k {
		case "mc":
			draw.MeshColor = v[0]
		case "mw":
			draw.MeshWidth = toFloat(v[0])
		case "cells":
			draw.Cells = toInt(v[0])
		}
	}
	draw.Print(w)
}

func toInt(s string) int {
	c, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("target=toInt, reason=strconv.Atoi, message=%s", err)
		return 1
	}
	return c
}

func toFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Printf("target=toFloat, reason=strconv.ParseFloat, message=%s", err)
		return 1
	}
	return f
}
