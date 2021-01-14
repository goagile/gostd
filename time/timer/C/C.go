package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.NewTimer(2 * time.Second)
	go func() {
		select {
		case <-a.C:
			fmt.Println("After 2 sec timer a stopped")
		}
	}()

	b := time.NewTimer(3 * time.Second)
	go func() {
		select {
		case <-b.C:
			fmt.Println("After 3 sec timer b stopped")
		}
	}()

	fmt.Scanln()
}
