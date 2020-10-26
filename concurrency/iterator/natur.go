package main

import (
	"fmt"
)

func main() {

	integers := make(chan int)

	// iterate over slice and put each i to chan
	go func() {
		for _, i := range []int{1, 2, 3, 4, 5} {
			integers <- i
		}
		close(integers)
	}()

	// iterate over chan and print each i
	for i := range integers {
		fmt.Println(i)
	}

}
