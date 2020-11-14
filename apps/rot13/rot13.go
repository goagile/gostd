package main

import (
	// "fmt"
	"io"
	"strings"
	"os"
	"unicode"
)

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := Rot13{s}
	io.Copy(os.Stdout, &r)
}

type Rot13 struct {
	r io.Reader
}

func (r13 Rot13) Read(out []byte) (int, error) {
	b := make([]byte, len(out))
	n, err := r13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := range out {
		out[i] = byte(rot13(rune(b[i])))
	}
	return n, nil
}


func rot13(s rune) rune {
	return rot(13, s)
}

func rot(step int, r rune) rune {
	s := unicode.ToLower(r)
	if s < 'a' || s > 'z' {
		return s
	}
	v := s + int32(step)
	if v > 'z' {
		d := v - 'z'
		v = 'a' + d - 1
	}
	return v
}
