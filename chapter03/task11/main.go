package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "123456789"
	fmt.Printf("s = %s, comma(s) = %s\n", s, comma(s))
	s = "-12345.6789"
	fmt.Printf("s = %s, comma(s) = %s\n", s, comma(s))
	s = "1234567.89"
	fmt.Printf("s = %s, comma(s) = %s\n", s, comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	if s_, found := strings.CutPrefix(s, "-"); found {
		buf.WriteString("-")
		s = s_
	}
	var n int
	if n = strings.Index(s, "."); n == -1 {
		n = len(s)
	}
	m := n % 3
	buf.WriteString(s[0:m])
	for i := m; i < n; i += 3 {
		if m > 0 || i != m {
			buf.WriteString(",")
		}
		buf.WriteString(s[i : i+3])
	}
	if n < len(s) {
		buf.WriteString(s[n:])
	}
	return buf.String()
}
