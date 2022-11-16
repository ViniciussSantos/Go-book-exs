package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordFreq := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()

		wordFreq[word]++
	}

	for word, n := range wordFreq {
		fmt.Printf("%s\t%d\n", word, n)
	}

}
