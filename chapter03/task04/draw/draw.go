package draw

import (
	"fmt"
	"io"
	"math"
)

var (
	Width, Height = 600, 320                     // Размер канвы в пикселях
	Cells         = 100                          // Количество ячеек сетки
	XYRange       = float64(30.0)                // Диапазон осей (-xyrange..+ xyrange)
	XYScale       = float64(Width) / 2 / XYRange // Пикселей в единице х или у
	ZScale        = float64(Height) * 0.4        // Пикселей в единице z
	angle         = math.Pi / 6                  // Углы осей х, у (=30°)
	minZ, maxZ    = -1., 1.
	MeshColor     = "grey"
	MeshWidth     = 0.5
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30°),cos(30°)

func Print(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: %f' "+
		"width='%d' height='%d' >\n", MeshColor, MeshWidth, Width, Height)
	for i := 0; i < Cells; i++ {
		for j := 0; j < Cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color(i, j))
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	// Ищем угловую точку (x,y) ячейки (i,j).
	x := XYRange * (float64(i)/float64(Cells) - 0.5)
	y := XYRange * (float64(j)/float64(Cells) - 0.5)
	// Вычисляем высоту поверхности z
	z := f(x, y)
	// Изометрически проецируем (x,y,z) на двумерную канву SVG (sx,sy)
	sx := float64(Width)/2 + (x-y)*cos30*XYScale
	sy := float64(Height)/2 + (x+y)*sin30*XYScale - z*ZScale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // Расстояние от (0,0)
	if r == 0.0 {
		return 0.0
	}
	return math.Sin(r) / r
}

func color(i, j int) uint32 {
	// Ищем угловую точку (x,y) ячейки (i,j)
	x := XYRange * (float64(i)/float64(Cells) - 0.5)
	y := XYRange * (float64(j)/float64(Cells) - 0.5)

	c := uint32((f(x, y) - minZ) / (maxZ - minZ) * 256)
	return (c << 16) | (c ^ 0xff)
}
