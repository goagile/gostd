package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 << 1, 1 >> 1) // 2
	fmt.Println(1 << 2, 2 >> 1) // 4 
	fmt.Println(1 << 3, 3 >> 1) // 8
	fmt.Println(1 << 4, 4 >> 1) // 16
	fmt.Println(1 << 5, 5 >> 1) // 16
	fmt.Println(1 << 6, 6 >> 1) // 16
}
