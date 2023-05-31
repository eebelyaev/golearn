package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	w := "go"
	fmt.Println("word: ", w)
	sha := flag.Int("sha", 256, "SHAxxx")
	flag.Parse()
	print(*sha, []byte(w))
}

func print(n int, w []byte) {
	switch n {
	case 384:
		fmt.Printf("digest (384): %x\n", sha512.Sum384([]byte(w)))
	case 512:
		fmt.Printf("digest (512): %x\n", sha512.Sum512([]byte(w)))
	default:
		fmt.Printf("digest (256): %x\n", sha256.Sum256([]byte(w)))
	}
}
