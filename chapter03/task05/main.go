/*
Упражнение 3.5. Реализуйте полноцветное множество Мандельброта с
использованием функции image.NewRGBA и типа color.RGBА или color.YCbCr.
*/
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
)

func main() {
	file, err := os.Create("m.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	print(file)
}

func print(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Точка (px, py) представляет комплексное значение z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // Примечание: игнорируем ошибки
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			c := 0xff - contrast*n
			switch n % 3 {
			case 0:
				return color.RGBA{c, n, n, 0xff}
			case 1:
				return color.RGBA{n, c, n, 0xff}
			case 2:
				return color.RGBA{n, n, c, 0xff}
			}
			return color.RGBA{c, 0x20, 0x77, 0xff}
		}
	}
	return color.White
}
