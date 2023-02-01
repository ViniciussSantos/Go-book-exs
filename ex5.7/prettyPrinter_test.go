package main

import (
	"testing"
)

func TestPrettyOutputCanBeParsed(t *testing.T) {
	// Given
	url := "https://golang.org"

	err := Outline(url)
	if err != nil {
		t.Error(err)
	}

}
