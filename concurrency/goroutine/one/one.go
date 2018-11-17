package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello 1")
	}()

	time.Sleep(1 * time.Millisecond)
}
