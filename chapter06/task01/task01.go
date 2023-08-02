package intset

// Len Возвращает количество элементов.
func (s *IntSet) Len() (res int) {
	for _, w := range s.words {
		for i := 0; i < 6; i++ {
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
