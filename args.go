package main

import (
	"fmt"
	"os"
)


func main() {

	var s string
	var n int

	n = len(os.Args)

	for i := 0; i < n; i++ {
		s += os.Args[i] + "\n"
	}

	fmt.Printf(s)

}
