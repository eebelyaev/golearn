package minmax

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		input   []int
		waitRes int
		waitOk  bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{}, 0, false},
		{[]int{3, 2, 1}, 1, true},
		{[]int{1, 1, 1}, 1, true},
	}
	for _, test := range tests {
		res, ok := Min(test.input...)
		if ok != test.waitOk || res != test.waitRes {
			t.Errorf("wait: ok = %t, res = %d, got: ok = %t, res = %d",
				test.waitOk, test.waitRes, ok, res)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input   []int
		waitRes int
		waitOk  bool
	}{
		{[]int{1, 2, 3}, 3, true},
		{[]int{}, 0, false},
		{[]int{3, 2, 1}, 3, true},
		{[]int{1, 1, 1}, 1, true},
	}
	for _, test := range tests {
		res, ok := Max(test.input...)
		if ok != test.waitOk || res != test.waitRes {
			t.Errorf("wait: ok = %t, res = %d, got: ok = %t, res = %d",
				test.waitOk, test.waitRes, ok, res)
		}
	}
}
