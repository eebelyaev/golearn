package intset

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
