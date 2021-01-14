package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339))

	go func() {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println(time.Now().Format(time.RFC3339))
		}
	}()

	fmt.Scanln()
	fmt.Println("ok")
}
