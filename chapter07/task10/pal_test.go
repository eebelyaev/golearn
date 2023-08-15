package pal

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input []string
		wait  bool
	}{
		{[]string{}, true},
		{[]string{"1", "2"}, false},
		{[]string{"1", "2", "3"}, false},
		{[]string{"1", "2", "2", "1"}, true},
		{[]string{"1", "2", "3", "2", "1"}, true},
	}
	for _, test := range tests {
		got := IsPalindrome(sort.StringSlice(test.input))
		if test.wait != got {
			t.Errorf("input = %q: wait = %t, got = %t", test.input, test.wait, got)
		}
	}
}
