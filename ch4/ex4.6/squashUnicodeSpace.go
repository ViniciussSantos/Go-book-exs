package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := []rune("teste     teste") // five spaces

	s = squashUnicodeSpace(s)
	fmt.Println(string(s))

}

func squashUnicodeSpace(s []rune) []rune {

	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(s[i]) && unicode.IsSpace(s[i+1]) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
