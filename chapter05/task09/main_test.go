package main

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		input, key string
		want       string
	}{
		{"12345", "", "12345"},
		{"123$foo45", "foo", fmt.Sprintf("123%s45", getVal("foo"))},
		{"123$foo45", "f", fmt.Sprintf("123%soo45", getVal("f"))},
		{"1$2345", "234", fmt.Sprintf("1%s5", getVal("234"))},
		{"w$elfkwg", "elf", fmt.Sprintf("w%skwg", getVal("elf"))},
		{"w$elfkwg$elfkw$elfkw", "elf",
			fmt.Sprintf("w%skwg%skw%skw",
				getVal("elf"), getVal("elf"), getVal("elf"))},
	}

	for _, test := range tests {
		key = test.key
		got := expand(test.input, getVal)
		if got != test.want {
			t.Errorf("expand(\"%s\", f(\"%s\")) = \"%s\", требуется \"%s\"",
				test.input, key, got, test.want)
		}
	}
}
