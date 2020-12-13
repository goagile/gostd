package main

import (
	"fmt"

	"github.com/goagile/gostd/bits/bit"
)

func main() {
	x := bit.Bit(5)

	// 00000101
	fmt.Printf("x = %08b\n", x)
	fmt.Println()

	fmt.Printf("x.Get(0) -> %v\n", x.Get(0))
	fmt.Printf("x.Get(1) -> %v\n", x.Get(1))
	fmt.Printf("x.Get(2) -> %v\n", x.Get(2))
	fmt.Printf("x.Get(3) -> %v\n", x.Get(3))
	fmt.Println()

	y := x.Set(1)
	// 00000111
	fmt.Printf("y = x.Set(1) = %08b\n", y)
	fmt.Printf("y.Get(1) -> %v\n", y.Get(1))
	fmt.Println()

	z := y.Clear(0)
	// 00000110
	fmt.Printf("z = y.Clear(0) = %08b\n", z)
	fmt.Printf("z.Get(0) -> %v\n", z.Get(0))
	fmt.Println()

	u := z.Update(1, false)
	// 00000100
	fmt.Printf("u := z.Update(1, false) = %08b\n", u)
	fmt.Printf("u.Get(1) -> %v\n", u.Get(1))
	fmt.Println()

	q := u.Update(3, true)
	// 00001100
	fmt.Printf("q := u.Update(3, true) = %08b\n", q)
	fmt.Printf("q.Get(3) -> %v\n", q.Get(3))
	fmt.Println()
}
