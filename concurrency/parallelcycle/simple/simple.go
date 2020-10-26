package main

import (
	"fmt"
	"time"
)

func main() {

	sqrs := make(chan int)

	n := 50

	for i := 0; i < n; i++ {
		go func(i int) {
			x := i * i
			time.Sleep(time.Duration(x) * time.Millisecond)
			sqrs <- x
		}(i)
	}

	for i := 0; i < n; i++ {
		fmt.Print(<-sqrs, " ")
	}
	fmt.Println()
}
