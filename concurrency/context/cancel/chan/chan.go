package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan struct{})

	go work(cancel)

	fmt.Scanln()

	cancel <- struct{}{}

	time.Sleep(1 * time.Second)
}

func work(cancel chan struct{}) {
	for {
		select {
		default:
			fmt.Print("work ")
			time.Sleep(100 * time.Millisecond)
		case <-cancel:
			fmt.Println("work Done")
			return
		}
	}
}
