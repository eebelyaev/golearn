// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import (
	"testing"
)

// !+String
func TestString(t *testing.T) {
	tests := []struct {
		expr string
	}{
		{"sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)"},
		{"5 / 9 * (F - 32)"},
		{"-1 + -x"},
		{"-1.12345 - x"},
		{"a / (b * c)"},
		{"a / b * c"},
		{"a + b - c"},
		{"a - (b - c)"},
		{"pow(x, 3 - pi) + sqrt(-(y + x) - (-2 / 3 * y + (14 - 10) * (x - 1)))"},
	}
	var es string
	for _, test := range tests {
		if expr, err := Parse(test.expr); err == nil {
			es = expr.String()
			if test.expr != es {
				t.Errorf("%s != %s\n", test.expr, es)
			}
		}
	}
}

//!-String
