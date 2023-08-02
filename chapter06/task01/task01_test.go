package intset

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	tests := [][]int{{}, {0}, {1, 2}, {1, 2, 60}, {1, 64, 2047, 2048}}
	for _, input := range tests {
		var s IntSet
		for _, v := range input {
			s.Add(v)
		}
		if s.Len() != len(input) {
			t.Errorf("%v.Len: %d, wait: %d", input, s.Len(), len(input))
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		input []int
		num   int
		wait  []int
	}{
		{[]int{0, 64, 128}, 3, []int{0, 64, 128}},
		{[]int{0, 64, 128}, 64, []int{0, 128}},
		{[]int{}, 0, []int{}},
	}
	for i, test := range tests {
		var wait IntSet
		for _, v := range test.wait {
			wait.Add(v)
		}
		var got IntSet
		for _, v := range test.input {
			got.Add(v)
		}
		got.Remove(test.num)
		if err := compare(got, wait); err != nil {
			t.Errorf("%d. wait: %v, got: %v", i+1, &wait, &got)
		}
	}
}

func TestClear(t *testing.T) {
	s := IntSet{[]uint64{0b1, 0b1, 0b1}}
	s.Clear()
	if len(s.words) != 0 {
		t.Errorf("множество не пустое: %v", s)
	}
}

func TestCopy(t *testing.T) {
	var s IntSet
	for _, v := range []int{0, 1, 128} {
		s.Add(v)
	}
	cp := *s.Copy()
	if err := compare(s, cp); err != nil {
		t.Error(err)
	}
	s.Add(5)
	if err := compare(s, cp); err == nil {
		t.Errorf("множества должны отличаться: %v !=  %v", &s, &cp)
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
