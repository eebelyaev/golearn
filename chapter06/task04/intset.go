package intset

import (
	"bytes"
	"fmt"
)

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint64
}

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество,
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len Возвращает количество элементов.
func (s *IntSet) Len() (res int) {
	for _, w := range s.words {
		for i := 0; i < 64; i++ {
			res += int(w & (1 << i) >> i)
		}
	}
	return res
}

// Remove Удаляет x из множества.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

// Clear Удаляет все элементы множества.
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy Возвращает копию множества.
func (s *IntSet) Copy() *IntSet {
	var cp IntSet
	if s.words != nil {
		cp.words = append(make([]uint64, 0), s.words...)
	}
	return &cp
}

// Compare сравнивает два множества.
func Compare(s1, s2 IntSet) error {
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

// AddAll позволяет добавлять список значений.
func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// IntersectWith - пересечение множеств.
func (s *IntSet) IntersectWith(s2 *IntSet) {
	if len(s.words) > len(s2.words) {
		s.words = s.words[:len(s2.words)]
	}
	for i := len(s.words) - 1; i >= 0; i-- {
		s.words[i] &= s2.words[i]
		if s.words[i] == 0 && len(s.words)-1 == i {
			s.words = s.words[:i]
		}
	}
}

// DifferenceWith вычитает переданное множество из множества-получателя.
func (s *IntSet) DifferenceWith(s2 *IntSet) {
	for i := len(s.words) - 1; i >= 0; i-- {
		if i < len(s2.words) {
			s.words[i] &^= s2.words[i]
			if s.words[i] == 0 && len(s.words)-1 == i {
				s.words = s.words[:i]
			}
		}
	}
}

// SymmetricDifference - Симметричная разность двух множеств содержит
// элементы, имеющиеся в одном из множеств, но не в обоих одновременно.
func (s *IntSet) SymmetricDifference(s2 *IntSet) {
	cut := true
	minlen := len(s2.words)
	if minlen > len(s.words) {
		minlen = len(s.words)
		s.words = append(s.words, s2.words[minlen:]...)
		cut = false
	}
	for i := minlen - 1; i >= 0; i-- {
		s.words[i] ^= s2.words[i]
		if cut && s.words[i] == 0 && len(s.words)-1 == i {
			s.words = s.words[:i]
		}
	}
	// cs := s.Copy()
	// s.DifferenceWith(s2)
	// cs2 := s2.Copy()
	// cs2.DifferenceWith(cs)
	// s.UnionWith(cs2)
}
