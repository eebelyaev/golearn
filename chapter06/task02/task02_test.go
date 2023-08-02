package intset

import (
	"fmt"
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
		if err := compare(wait, got); err != nil {
			t.Errorf("%d. wait: %v, got: %v", i+1, wait, got)
		}
	}
}

func compare(s1, s2 IntSet) error {
	if s1.Len() != s2.Len() {
		return fmt.Errorf("отличается длина: s1.Len = %d, s2.Len = %d",
			s1.Len(), s2.Len())
	}
	for i, w := range s1.words {
		if s2.words[i] != w {
			return fmt.Errorf("элементы слайсов words[%d] отличаются: %d != %d",
				i, w, s2.words[i])
		}
	}
	return nil
}
