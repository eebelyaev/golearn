/*
Упражнение 3.4. Создайте веб-сервер, который вычисляет поверхности и
возвращает клиенту SVG-данные.
*/
package main

import (
	"fmt"
	cartab "golearn/chapter07/task09/cartab"
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"
)

var templ = `
<h1>List of Cars</h1>
<table>
	<tr style='text-align: left'>
		<th><a href="/?col=0">Brand</a></th>
		<th><a href="/?col=1">Model</a></th>
		<th><a href="/?col=2">Year</a></th>
	</tr>
	{{range .}}
	<tr>
		<td>{{.Brand}}</td>
		<td>{{.Model}}</td>
		<td>{{.Year}}</td>
	</tr>
	{{end}}
</table>
<p>{{printStateSort}}</p>
`
var report = template.Must(template.New("carlist").
	Funcs(template.FuncMap{"printStateSort": printStateSort}).
	Parse(templ))
var stateSorts = [3]cartab.StateSort{
	{Act: false, Col: 0, Asc: true},
	{Act: false, Col: 1, Asc: true},
	{Act: false, Col: 2, Asc: true},
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if v, exist := r.URL.Query()["col"]; exist {
		col, err := strconv.Atoi(v[0])
		if err != nil || col < 0 || col >= len(stateSorts) {
			log.Printf("error: %v, col = %s\n", err, v[0])
		} else {
			click(col)
			sort.Sort(cartab.SortCars{Cars: cartab.Cars, Orders: stateSorts})
		}
	}
	if err := report.Execute(w, cartab.Cars); err != nil {
		log.Fatal(err)
	}
}

func click(col int) {
	ss := stateSorts[:]
	if !ss[0].Act {
		ss[0] = cartab.StateSort{Act: true, Col: col, Asc: true}
	} else if ss[0].Col == col {
		ss[0].Asc = !ss[0].Asc
	} else {
		for i, s := range ss[1:] {
			if s.Col == col {
				copy(ss[1:i+2], ss[0:i+1])
				break
			}
		}
		ss[0].Col, ss[0].Asc = col, true
	}
}

func printStateSort() string {
	return fmt.Sprintf("%v", stateSorts)
}
