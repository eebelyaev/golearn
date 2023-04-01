package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	lissajousPng()
}

func drawCircle() {
	// Создаем новое изображение размером 200x200 пикселей
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))

	// Заполняем изображение белым цветом
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// Определяем центр круга и его радиус
	x := 100
	y := 100
	r := 50

	// Рисуем круг на изображении
	for i := 0; i < 360; i++ {
		// Вычисляем координаты точки на окружности
		rad := float64(i) * math.Pi / 180.0
		px := float64(x) + float64(r)*math.Cos(rad)
		py := float64(y) + float64(r)*math.Sin(rad)

		// Устанавливаем цвет точки на изображении
		img.Set(int(px), int(py), color.Black)
	}

	// Сохраняем изображение в файл
	file, err := os.Create("circle.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Круг успешно нарисован и сохранен в файле circle.png")
}

func lissajousPng() {
	const (
		cycles  = 10
		res     = 0.001
		size    = 500
		nframes = 64
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0
	rect := image.Rect(0, 0, 2*size+1, 2*size+1)
	img := image.NewPaletted(rect, []color.Color{color.White, color.Black})
	phase := 0.0
	for i := 0; i < nframes; i++ {
		for t := 0.0; t < cycles*math.Pi*2; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
	}

	fileName := "lissajous.png"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		file, err = os.Create("lissajous.png")
		if err != nil {
			panic(err)
		}
	}
	// Сохраняем изображение в файл
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Изображение успешно нарисовано и сохранено в файл lissajous.png")
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, []color.Color{color.White, color.Black})
		for t := 0.0; t < cycles*math.Pi*2; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 0)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
