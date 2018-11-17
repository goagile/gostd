package main

import "fmt"

type result chan int

func main() {
	results := []result{}

	// create future results
	for i := 0; i < 10; i++ {
		r := sum(i, 1000)
		results = append(results, r)
	}

	// calculate results
	for _, r := range results {
		fmt.Println(<-r)
	}
}

func sum(a, b int) result {
	r := make(result)
	go func() {
		r <- a + b
	}()
	return r
}
