package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		input IntSet
		add   []int
	}{
		{IntSet{}, []int{0}},
		{IntSet{[]uint64{0b1}}, []int{2}},
		{IntSet{[]uint64{0b1}}, []int{0, 2, 4, 128}},
	}

	for i, test := range tests {
		wait := *test.input.Copy()
		for _, a := range test.add {
			wait.Add(a)
		}
		got := *test.input.Copy()
		got.AddAll(test.add...)
		if err := Compare(wait, got); err != nil {
			t.Errorf("%d. wait: %v, got: %v", i+1, wait, got)
		}
	}
}
