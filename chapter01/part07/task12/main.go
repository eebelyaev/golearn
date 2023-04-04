package main

import (
	"fmt"
	"golearn/chapter01/part07/task12/gifgen"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var g gifgen.Gif

func main() {
	g = gifgen.New()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/gif")
	// for i, q := range r.URL.Query() {
	// 	fmt.Fprintf(os.Stdout, "%q\t%q", i, q)
	// }
	if r.URL.Query().Has("cycles") {
		c, err := strconv.Atoi(r.URL.Query().Get("cycles"))
		if err != nil {
			log.Printf("from: handler(http.ResponseWriter, *http.Request), error: %s", err)
			return
		}
		g.Cycles = c
	}
	lissajous(w)
}

func lissajous(w io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: g.Nframes}
	phase := 0.0
	fmt.Fprintf(os.Stdout, "%f\n", freq)
	for i := 0; i < g.Nframes; i++ {
		rect := image.Rect(0, 0, 2*g.Size+1, 2*g.Size+1)
		img := image.NewPaletted(rect, []color.Color{color.White, color.Black})
		for t := 0.0; t < float64(g.Cycles)*math.Pi*2; t += g.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(g.Size+int(x*float64(g.Size)+0.5), g.Size+int(y*float64(g.Size)+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, g.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
