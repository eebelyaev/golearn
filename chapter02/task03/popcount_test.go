package main

import (
	pc0 "golearn/chapter02/popcount"
	pc "golearn/chapter02/task03/popcount"
	"testing"
)

func BenchmarkPopCount0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc0.PopCount(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.PopCount(uint64(i))
	}
}
