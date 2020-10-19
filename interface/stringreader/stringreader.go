package main

import (
	"fmt"
	"io"
)

func main() {
	s := NewReader("he__ll__o!__")
	for {
		b := make([]byte, 2)
		_, err := s.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Print(string(b), "\n")
	}
}

func NewReader(s string) *reader {
	return &reader{s: s}
}

type reader struct {
	s string
	i int
}

func (r *reader) Read(p []byte) (n int, err error) {
	if r.i >= len(r.s) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += n
	return
}
