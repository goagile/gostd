package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello 1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello 2")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello 3")
	}()

	wg.Wait()
}
