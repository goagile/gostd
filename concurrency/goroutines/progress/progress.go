package main

import (
	"fmt"
	"time"
)

func main() {

	nums := make(chan int)

	go progress(nums)

	for i := 1; i <= 100; i++ {
		time.Sleep(100 * time.Millisecond)
		nums <- i
	}

	fmt.Printf("\r          \n")
}

func progress(nums chan int) {
	for i := range nums {
		fmt.Printf("\r%02v", i)
	}
}
