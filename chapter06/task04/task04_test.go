package intset

import (
	"testing"
)

func TestElems(t *testing.T) {
	tests := [][]int{{}, {1}, {1, 2, 16, 256, 1024}}
	for _, elems := range tests {
		var s, s2 IntSet
		s.AddAll(elems...)
		s2.AddAll(s.Elems()...)
		if err := Compare(s, s2); err != nil {
			t.Error(err)
		}
	}
}
