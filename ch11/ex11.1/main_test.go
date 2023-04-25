package charcount

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {

	var tests = []struct {
		input  string
		counts map[rune]int
		utflen [utf8.UTFMax + 1]int
	}{
		{"", map[rune]int{}, [utf8.UTFMax + 1]int{}},
		{"a", map[rune]int{'a': 1}, [utf8.UTFMax + 1]int{0, 1, 0, 0, 0}},
		{"aa", map[rune]int{'a': 2}, [utf8.UTFMax + 1]int{0, 2, 0, 0, 0}},
		{"ab", map[rune]int{'a': 1, 'b': 1}, [utf8.UTFMax + 1]int{0, 2, 0, 0, 0}},
		{"kayak", map[rune]int{'k': 2, 'a': 2, 'y': 1}, [utf8.UTFMax + 1]int{0, 5, 0, 0, 0}},
		{"detartrated", map[rune]int{'d': 2, 'e': 2, 't': 3, 'a': 2, 'r': 2}, [utf8.UTFMax + 1]int{0, 11, 0, 0, 0}},
	}

	for _, test := range tests {
		counts, utflen := charcount(strings.NewReader(test.input))
		if !reflect.DeepEqual(counts, test.counts) || !reflect.DeepEqual(utflen, test.utflen) {
			t.Errorf("charcount(%q) = %v, %v; want %v, %v", test.input, counts, utflen, test.counts, test.utflen)
		}
	}
}
