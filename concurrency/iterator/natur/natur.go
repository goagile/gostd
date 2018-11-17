package main

import (
	"fmt"
	"time"
)

func main() {

	integers := make(chan int)

	// iterate over slice and put each i to chan
	go func() {
		for _, i := range []int{1, 2, 3, 4, 5} {
			integers <- i
		}
	}()

	// iterate over chan and print each i
	go func() {
		for i := range integers {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second)

	close(integers)

}
