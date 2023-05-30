// comma вставляет запятые в строковое представление
// неотрицательного десятичного числа.
package main

import "fmt"

func main() {
	s := "1234567"
	fmt.Printf("s = %s, comma(s) = %s", s, comma(s))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
