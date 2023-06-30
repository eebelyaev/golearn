package main

import (
	"fmt"
	"strings"
)

var key = "foo"

func main() {
	s := "football $foo"
	res := expand(s, getVal)
	fmt.Printf("s = %s, expand s = %s", s, res)
}

func expand(s string, f func(string) string) string {
	return strings.ReplaceAll(s, "$"+key, f(key))
}

func getVal(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("~[%s]~", s)
}
