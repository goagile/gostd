package main

import (
	"fmt"
)

func main() {
	ints := make(chan int)
	sqrs := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ints <- i
		}
		close(ints)
	}()

	go func() {
		for i := range ints {
			sqrs <- i * i
		}
		close(sqrs)
	}()

	for s := range sqrs {
		fmt.Printf("%v ", s)
	}
	fmt.Println()
}
