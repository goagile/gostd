package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(1 * time.Millisecond)
	var i int
	for {
		select {
		default:
			fmt.Print(i, " ")
			i++
		case <-t.C:
			goto AFTER
		}
	}
AFTER:
	fmt.Println("OK")
}
