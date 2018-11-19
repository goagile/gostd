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
			sqrs <- square(i)
		}(i)
	}

	for i := 0; i < n; i++ {
		fmt.Println(<-sqrs)
	}
}

func square(i int) int {
	x := i * i
	time.Sleep(time.Duration(x) * time.Millisecond)
	return x
}
