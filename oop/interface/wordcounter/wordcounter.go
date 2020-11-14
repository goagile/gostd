package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "  aa bbb ccc, hello, world! "
	var c WordCounter
	fmt.Fprintf(&c, "%v", s)
	fmt.Println(s, "have", c, "words")
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c = WordCounter(len(strings.Fields(string(p))))
	return int(*c), nil
}
