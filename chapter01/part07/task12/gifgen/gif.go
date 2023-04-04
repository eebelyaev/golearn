package gifgen

import "image/gif"

type Gif struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
	Gif     gif.GIF
}

func New() Gif {
	g := gif.GIF{LoopCount: 64}
	return Gif{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		Nframes: 64,
		Delay:   8,
		Gif:     g,
	}
}
