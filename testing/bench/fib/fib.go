package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, ": ", fibRecur(i))
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i, ": ", fibCicle(i))
	}
}

func fibRecur(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibRecur(n-1) + fibRecur(n-2)
}

func fibCicle(n int) int {
	a, b := 0, 1
	for b <= n {
		a, b = b, a+b
	}
	return b
}
