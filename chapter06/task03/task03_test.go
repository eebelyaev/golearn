package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		s, s2, wait []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{}, []int{}},
		{[]int{1, 2, 3, 128}, []int{3, 16, 128}, []int{3, 128}},
		{[]int{1, 2, 3, 128}, []int{128, 256}, []int{128}},
		{[]int{1, 2, 3, 128, 1024}, []int{3, 16, 256}, []int{3}},
	}
	for _, test := range tests {
		var s, s2, wait IntSet
		s.AddAll(test.s...)
		s2.AddAll(test.s2...)
		wait.AddAll(test.wait...)
		got := *s.Copy()
		got.IntersectWith(&s2)
		if err := Compare(wait, got); err != nil {
			t.Errorf("%v.f(%v) = %v, wait: %v", &s, &s2, &got, &wait)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		s, s2, wait []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1, 2, 3, 128}, []int{3, 16, 128}, []int{1, 2}},
		{[]int{1, 2, 3, 128}, []int{3, 16, 256, 1024}, []int{1, 2, 128}},
	}
	for _, test := range tests {
		var s, s2, wait IntSet
		s.AddAll(test.s...)
		s2.AddAll(test.s2...)
		wait.AddAll(test.wait...)
		got := *s.Copy()
		got.DifferenceWith(&s2)
		if err := Compare(wait, got); err != nil {
			t.Errorf("%v.f(%v) = %v, wait: %v", &s, &s2, &got, &wait)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		s, s2, wait []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{1, 2, 3, 128}, []int{3, 16, 128}, []int{1, 2, 16}},
		{[]int{1, 2, 3, 128}, []int{128, 256}, []int{1, 2, 3, 256}},
		{[]int{1, 2, 3, 1024}, []int{3, 16, 256}, []int{1, 2, 16, 256, 1024}},
	}
	for _, test := range tests {
		var s, s2, wait IntSet
		s.AddAll(test.s...)
		s2.AddAll(test.s2...)
		wait.AddAll(test.wait...)
		got := *s.Copy()
		got.SymmetricDifference(&s2)
		if err := Compare(wait, got); err != nil {
			t.Errorf("%v.f(%v) = %v, wait: %v", &s, &s2, &got, &wait)
		}
	}
}
