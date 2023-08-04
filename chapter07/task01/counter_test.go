package counter

import (
	"fmt"
	"testing"
)

func TestWordCounter(t *testing.T) {
	tests := []struct {
		input string
		wait  int
	}{
		{"Hello, world!", 2},
		{"1 2 3", 3},
		{"0", 1},
		{"", 0},
	}
	for _, test := range tests {
		var c WordCounter
		fmt.Fprint(&c, test.input)
		if c != WordCounter(test.wait) {
			t.Errorf("input: %s, wait: %d, got: %d", test.input, test.wait, c)
		}
	}
}

func TestLineCounter(t *testing.T) {
	tests := []struct {
		input string
		wait  int
	}{
		{`line 1
	line 2
	line 3`, 3},
		{"1\n2\n3", 3},
		{"0", 1},
		{"", 0},
	}
	for _, test := range tests {
		var c LineCounter
		fmt.Fprint(&c, test.input)
		if c != LineCounter(test.wait) {
			t.Errorf("input: %s, wait: %d, got: %d", test.input, test.wait, c)
		}
	}
}

func TestLineCounterPrintln(t *testing.T) {
	wait := 2
	var c LineCounter
	fmt.Fprintln(&c, "Доброе")
	fmt.Fprintln(&c, "утро")
	if c != LineCounter(wait) {
		t.Errorf("1. wait: %d, got: %d", wait, c)
	}
	wait = 3
	fmt.Fprintln(&c, "мир!")
	if c != LineCounter(wait) {
		t.Errorf("2. wait: %d, got: %d", wait, c)
	}
}
