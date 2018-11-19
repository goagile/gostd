package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	results := make(chan int)

	n := 10

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results <- square(i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for x := range results {
		fmt.Println(x)
	}

}

func square(i int) int {
	x := i * i
	time.Sleep(time.Duration(x) * time.Millisecond)
	return x
}
