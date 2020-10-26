package main

import (
	"fmt"
)

func main() {

	nums := make(chan int, 50)
	
	go genfib(nums)

	for n := range nums {
		fmt.Print(n, " ")
	}
	fmt.Println()

}

func genfib(nums chan<- int) {
	x, y := 1, 1
	for i := 0; i < cap(nums); i++ {
		nums <- x
		x, y = y, x+y
	}
	close(nums)
}
