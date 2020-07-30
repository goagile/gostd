package main

import (
	"fmt"
)

func main() {
	s := 0
	for x := 1; x <= 3; x ++ {
		s += x
	}
	fmt.Println(s)
}
