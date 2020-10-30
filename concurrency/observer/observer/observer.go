package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeout := time.After(4 * time.Second)
	numbers := make(chan int)
	done := make(chan struct{})
	go func() {
		var i int
		for {
			select {
			case <-done:
				close(numbers)
				return
			case numbers <- i:
				i++
			}
		}
	}()
	for {
		select {
		case <-timeout:
			done <- struct{}{}
			fmt.Println("\rOK ->")
			os.Exit(0)
		case i := <-numbers:
			fmt.Printf("\r  ... %v%%   ", i)
			time.Sleep(200 * time.Millisecond)
		}
	}
}
