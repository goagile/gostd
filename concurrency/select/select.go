package main

import (
	"fmt"
)

func main() {

	n := 10

	alpha := make(chan int, n)
	betta := make(chan int, n)

	go func() {
		for i := 1; i <= n; i++ {
			alpha <- i
			betta <- i
		}
		close(alpha)
		close(betta)
	}()

	for i := 0; i < n; i++ {
		select {
		case a := <-alpha:
			fmt.Printf("a:%v ", a)
		case b := <-betta:
			fmt.Printf("b:%v ", b)
		}
	}

}
