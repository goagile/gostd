package main

import (
	"fmt"
	"strings"
	"io"
)

func main() {
	r := strings.NewReader("Hello, world!")	
	
	var err error
	var n int
	b := make([]byte, 4)
	for err != io.EOF {
		n, err = r.Read(b)
		fmt.Printf("Read: %v bytes: %v %q\n", n, b[:n], b[:n])
	}

	fmt.Println("stop")
}