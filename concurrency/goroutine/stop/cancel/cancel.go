package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	nums := make(chan int)
	timeout := time.After(2 * time.Second)
	done := make(chan struct{})

	go func() {
		var i int
		for {
			select {
			case nums <- i:
				i++
			case <-done:
				return
			}
		}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for {
		select {
		case i := <-nums:
			fmt.Print("--->")
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
		case <-timeout:
			fmt.Println("\n ... Time is out!")
			return
		case <-done:
			fmt.Println("... Key pressed!")
			return
		}
	}
}
