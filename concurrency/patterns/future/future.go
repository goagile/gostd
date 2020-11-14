package main

import "fmt"

func main() {
	results := []chan int{}

	for i := 0; i < 10; i++ {
		r := sum(i, 1000)
		results = append(results, r)
	}

	// calculate
	for _, r := range results {
		fmt.Println(<-r)
	}
}

func sum(a, b int) chan int {
	r := make(chan int)
	go func() {
		r <- a + b
	}()
	return r
}
