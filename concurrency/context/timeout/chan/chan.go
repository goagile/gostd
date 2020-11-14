package main

import (
	"fmt"
	"time"
)

func main() {

	timeout := time.After(1 * time.Second)

	go work(timeout)

	fmt.Scanln()

	fmt.Println("OK")

}

func work(timeout <-chan time.Time) {
	for {
		select {
		default:
			fmt.Print("A ")
			time.Sleep(100 * time.Millisecond)
		case <-timeout:
			fmt.Print("Done by timeout")
			return
		}
	}
}
