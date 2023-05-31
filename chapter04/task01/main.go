package main

import (
	"crypto/sha256"
	"fmt"
	"golearn/chapter02/popcount"
)

func main() {
	digest1 := sha256.Sum256([]byte("x"))
	digest2 := sha256.Sum256([]byte("X"))
	fmt.Printf("digest1 = %x\ndigest2 = %x\n", digest1, digest2)
	fmt.Println("CountDifferentBits = ", CountDifferentBits(digest1, digest2))
}

// CountDifferentBits возвращает количество различных битов в двух дайджестах SHA256
func CountDifferentBits(digest1 [32]byte, digest2 [32]byte) int {
	count := 0
	for i := 0; i < len(digest1); i++ {
		count += popcount.PopCount(uint64(digest1[i] ^ digest2[i]))
	}
	return count
}
