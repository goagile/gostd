package main

import "fmt"

func main() {
	f := fact()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}

func fact() func() int {
	a := 1
	p := 1
	return func() int {
		p = p * a
		a++
		return p 
	}
}
