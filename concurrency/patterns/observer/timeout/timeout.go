package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeout := time.After(1 * time.Second)
	var i int
	for {
		select {
		case <-timeout:
			fmt.Println("OK")
			os.Exit(0)
		default:
			fmt.Print(i, " ")
			i++
			time.Sleep(200 * time.Millisecond)
		}
	}
}
