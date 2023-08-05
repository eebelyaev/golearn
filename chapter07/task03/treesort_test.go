package treesort

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		input []int
		wait  string
	}{
		{[]int{2, 4, 5, 1, 3}, "[1 2 3 4 5]"},
		{[]int{2, 3, 1}, "[1 2 3]"},
		{[]int{0}, "[0]"},
	}
	for _, test := range tests {
		tr := new(tree)
		for _, v := range test.input {
			add(tr, v)
		}
		got := tr.String()
		if test.wait != got {
			t.Errorf("wait: %s, got: %s", test.wait, got)
		}
	}
}
