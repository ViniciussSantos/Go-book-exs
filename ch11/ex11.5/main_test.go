package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	var tests = []struct {
		input string
		sep   string
		want  int
	}{
		{"", ":", 1},
		{"a", ":", 1},
		{"a:b", ":", 2},
		{"a:b:c", ":", 3},
	}

	for _, test := range tests {
		words := strings.Split(test.input, test.sep)
		if got := len(words); got != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d",
				test.input, test.sep, got, test.want)
		}
	}

}
