package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

const (
	width  = 500
	height = 500
	cycles = 64
)

func main() {
	images := make([]*image.Paletted, cycles)
	delays := make([]int, cycles)

	for i := 0; i < cycles; i++ {
		images[i] = image.NewPaletted(image.Rect(0, 0, width, height), color.Palette{color.White, color.Black})
		for x := 0; x < width; x++ {
			sinValue := math.Sin(float64(x+i*5)*2*math.Pi/width)*127.0 + 128.0
			y := int(sinValue)
			images[i].Set(x, y, color.Black)
		}
		delays[i] = 8
	}

	file, err := os.Create("sinusoid.gif")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}
