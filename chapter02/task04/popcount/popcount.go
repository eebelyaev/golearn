package popcount

// PopCount возвращает степень заполнения
// (количество установленных битов) значения х.
func PopCount(x uint64) int {
	var r int
	for i := 0; i < 64; i++ {
		r += int((x >> i) & 1)
	}
	return r
}
