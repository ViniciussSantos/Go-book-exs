package main

import (
	"io"
)

type Reader struct {
	s string
	i int
}

func (r *Reader) Read(b []byte) (n int, err error) {

	if len(b) == 0 {
		return 0, nil
	}

	n = copy(b, r.s[r.i:])
	r.i += n

	if r.i >= int(len(r.s)) {
		return 0, io.EOF
	}
	return
}

func NewReader(s string) *Reader { return &Reader{s, 0} }
