package main

import (
	"fmt"
)

func main() {
	nums := make(chan int)
	sqrs := make(chan int)

	go numbser(nums)
	go squarer(sqrs, nums)

	printer(sqrs)

	fmt.Println()
}

func numbser(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Printf("%v ", x)
	}
}
