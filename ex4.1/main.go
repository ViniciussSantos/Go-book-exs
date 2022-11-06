package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	diffBits := countDiffBits(c1[:], c2[:])

	fmt.Println(diffBits)
	fmt.Println(diffBits == 125)

}

// https://en.wikipedia.org/wiki/Hamming_weight
func popCount(b byte) int {
	count := 0

	for ; b != 0; count++ {
		b &= b - 1
	}

	return count
}

func countDiffBits(c1, c2 []byte) int {
	count := 0
	for i := 0; i < len(c1) || i < len(c2); i++ {
		//XOR gets only the different bits but only works if len(c1) == len(c2)
		count += popCount(c1[i] ^ c2[i])
	}
	return count
}
