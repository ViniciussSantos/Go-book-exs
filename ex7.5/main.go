package main

import "io"

type LimitedReader struct {
	r io.Reader
	n int
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if len(p) > l.n {
		p = p[:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= n

	return
}
