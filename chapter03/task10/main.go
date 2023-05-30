package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "123456789"
	fmt.Printf("s = %s, comma(s) = %s", s, comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	m := n % 3
	buf.WriteString(s[0:m])
	for i := m; i < n; i += 3 {
		if m > 0 || i != m {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
