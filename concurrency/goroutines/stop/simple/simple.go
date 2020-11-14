package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	data := make(chan int)
	done := make(chan struct{})
	go func() {
		var i int
		for {
			data <- i
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
			i++
		}
	}()
	go func() {
		if _, err := fmt.Scanln(); err != nil {
			done <- struct{}{}
		}
	}()
	for {
		select {
		case d := <-data:
			fmt.Print(d, " ")
		case <-done:
			os.Exit(0)
		}
	}
}
