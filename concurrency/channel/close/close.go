package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 10)
	
	go fib(cap(c), c)

	for {
		p, ok := <- c 
		if !ok {
			break
		}
		fmt.Printf("%v ", p)
	}

	fmt.Println("\nfinish")

}

func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
