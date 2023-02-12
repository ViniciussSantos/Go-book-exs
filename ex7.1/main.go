package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type LineCounter int
type WordCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		*l++
	}

	return len(p), s.Err()

}

func (w *WordCounter) Write(p []byte) (int, error) {

	s := bufio.NewScanner(bytes.NewBuffer(p))

	s.Split(bufio.ScanWords)

	for s.Scan() {
		*w++
	}

	return len(p), s.Err()

}

func main() {
	var w WordCounter

	w.Write([]byte("hello world"))
	fmt.Println(w) // "5", = len("hello")

	var l LineCounter
	l.Write([]byte("hello world\n"))
	fmt.Println(l)
}
