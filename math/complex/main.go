package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := 12 + 23i
	fmt.Printf("var x type: %T, value: %v\n", x, x)
	fmt.Println(cmplx.Sqrt(1 + 1i))
}
