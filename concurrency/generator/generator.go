package main

import (
	"fmt"
)

func main() {
	n := genints()
	for i := 0; i < 10; i++ {
		fmt.Print(<-n, " ")
	}
	fmt.Println()
}

func genints() chan int {
	n := make(chan int)
	var i int
	go func(i int) {
		for {
			n <- i
			i++
		}
	}(i)
	return n
}
