package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	a, b := 0, 1
	for a + b < n {
		a, b = b, a + b
	}
	return b
}

func main() {
	fmt.Println(fib(10))
}
