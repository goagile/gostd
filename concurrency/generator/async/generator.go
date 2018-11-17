package main

import (
	"fmt"
)

func main() {
	for i := range genInts(10) {
		fmt.Println(i)
	}
}

func genInts(n int) chan int {
	result := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			result <- i
		}()
	}
	return result
}
