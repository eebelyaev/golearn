package intset

// AddAll позволяет добавлять список значений.
func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}
