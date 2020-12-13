package main

import "fmt"

func main() {

	fmt.Println("Bit manipulations")
	fmt.Println()

	fmt.Println("4 bits:")
	fmt.Printf(" 1 = %04b\n", 1)
	fmt.Printf("-1 = %04b\n", -1)

	fmt.Printf(" 2 = %04b\n", 2)
	fmt.Printf("-2 = %04b\n", -2)

	fmt.Printf(" 3 = %04b\n", 3)
	fmt.Printf("-3 = %04b\n", -3)

	fmt.Printf(" 4 = %04b\n", 4)
	fmt.Printf("-4 = %04b\n", -4)

	a := 3 // 0011
	b := 5 // 0101

	fmt.Println()
	fmt.Printf("a = %04b\n", a)
	fmt.Printf("b = %04b\n", b)
	fmt.Println()

	// 0011 & 0101 = 0001
	fmt.Printf("Bitwice AND: a & b = %04b & %04b = %04b\n", a, b, a&b)

	// 0011 | 0101 = 0111
	fmt.Printf("Bitwice OR:  a | b = %04b | %04b = %04b\n", a, b, a|b)

	// 0011 ^ 0101 = 0110
	fmt.Printf("Bitwice XOR: a ^ b = %04b ^ %04b = %04b\n", a, b, a^b)

	// ^0101 = 1010
	fmt.Printf("Bitwice NOT: ^b = ^%04b = %04b\n", b, uint8(^b))

	// 0011 &^ 0101 = 0010
	fmt.Printf("Bitclear AND NOT: a &^ b = %04b &^ %04b = %04b\n", a, b, a&^b)

	// 0011 << 2 = 1100
	fmt.Printf("Left Shift: a << 2 = %04b << 2 = %04b\n", a, a<<2)

	// 0101 >> 2 = 0001
	fmt.Printf("Right Shift: b >> 2 = %04b >> 2 = %04b\n", b, b>>2)
	fmt.Println()

	// shift
	fmt.Println("Shift:")
	for i := 0; i < 8; i++ {
		fmt.Printf("1 << %v = %08b\n", i, 1<<i)
	}

	fmt.Println()
	var x uint8 = 3 << 1 // 00000110
	fmt.Printf("x = %08b {1,2}\n", x)

	var y uint8 = 13 << 1 // 00000110
	fmt.Printf("y = %08b {1,3,4}\n", y)

	// Conjunction
	// 00000110 | 00011010 = 00011110 {1,2,3,4}
	fmt.Printf("x | y = %08b | %08b = %08b {1,2,3,4}\n", x, y, x|y)

	// Disjunction
	// 00000110 & 00011010 = 00000010 {1}
	fmt.Printf("x & y = %08b & %08b = %08b {1}\n", x, y, x&y)

	// Simmetric Diff
	// 00000110 ^ 00011010 = 00011100 {2}
	fmt.Printf("x ^ y = %08b ^ %08b = %08b {2}\n", x, y, x^y)

	// Diff
	// 00000110 &^ 00011010 = 00000100 {2}
	fmt.Printf("x &^ y = %08b &^ %08b = %08b {2}\n", x, y, x&^y)
}
