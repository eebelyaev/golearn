// Выводит URL и описание каждого комикса, соответствующего
// условию поиска, заданному в командной строке. Сайт: xkcd.com
package main

import (
	"os"
	"strings"
)

func main() {
	comics := newComics().load(3)
	param := strings.ToLower(strings.Join(os.Args[1:], " "))
	for i := 0; i < len(comics); i++ {
		if strings.Contains(strings.ToLower(comics[i].Title), param) {
			comics[i].print()
		}
	}
	comics.save()
}
