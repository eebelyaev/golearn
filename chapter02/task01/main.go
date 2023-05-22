/*
Упражнение 2.1. Добавьте в пакет tempconv типы, константы и функции
для работы с температурой по шкале Кельвина, в которой нуль градусов
соответствует температуре -273.15°С, а разница температур в 1К имеет
ту же величину, что и 1°С.
*/
package main

import (
	"fmt"
	"golearn/chapter02/task01/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		k := tempconv.Kelvin(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			k, tempconv.KToC(k), c, tempconv.CToK(c))
	}
}
