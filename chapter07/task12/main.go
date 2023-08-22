package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

var mu sync.Mutex
var templ = `
<h1>Items</h1>
<table>
	<tr style='text-align: left'>
		<th><a href="/?col=0">Item</a></th>
		<th><a href="/?col=1">Price</a></th>
	</tr>
	{{range $index, $element := .}}
	<tr>
		<td>{{$index}}</td>
		<td>{{$element}}</td>
	</tr>
	{{end}}
</table>
`
var report = template.Must(template.New("carlist").
	Parse(templ))

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := report.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "товар не указан\n")
		return
	}
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "ошибка распознавания цены: %q\n", err)
		return
	}
	mu.Lock()
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "товар уже существует: %q\n", item)
		return
	}
	db[item] = dollars(p)
	mu.Unlock()
	fmt.Fprintf(w, "создан товар %q за %s\n", item, db[item])
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "ошибка распознавания цены: %v\n", err)
		return
	}
	mu.Lock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	db[item] = dollars(p)
	mu.Unlock()
	fmt.Fprintf(w, "обновлена цена товара %q: %s\n", item, db[item])
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "нет товара: %q\n", item)
		return
	}
	delete(db, item)
	mu.Unlock()
	fmt.Fprintf(w, "товар %q удален\n", item)
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars
