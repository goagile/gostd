package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello 1")
	}()

	go func() {
		fmt.Println("Hello 2")
	}()

	go func() {
		fmt.Println("Hello 3")
	}()

	time.Sleep(1 * time.Millisecond)
}
