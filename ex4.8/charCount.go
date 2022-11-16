package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	unicodeTypes := make(map[string]int)

	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, nbytes, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && nbytes == 1 {
			invalid++
			continue
		}

		for propertyName, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				unicodeTypes[propertyName]++
			}
		}
		counts[r]++
		utflen[nbytes]++

	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nunicode type\tcount\n")

	for u, n := range unicodeTypes {
		fmt.Printf("%q\t%d\n", u, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
