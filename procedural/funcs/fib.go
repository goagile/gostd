package main

import (
	"fmt"
)

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fib() func() int {
	a := 0 
	b := 1
	return func() int {
		c := a + b
		b = a
		a = c
		return c
	}
}

