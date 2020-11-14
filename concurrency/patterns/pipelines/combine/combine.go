package main

import "fmt"

func main() {

	nums := []int{}
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}

	print(
		sqrt(
			list(nums)))

}

func print(in <-chan int) {
	for x := range in {
		fmt.Printf("%v ", x)
	}
}

func sqrt(in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()
	return out
}

func list(in []int) chan int {
	out := make(chan int)
	go func() {
		for _, x := range in {
			out <- x
		}
		close(out)
	}()
	return out
}
