package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	kill := make(chan struct{})

	// broadcast
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(kill)
	}()

	ws := "ABC"

	// workers
	for i, name := range ws {
		go func(i int, name rune) {
			for {
				select {
				default:
					fmt.Print(string(name), " -> ")
					time.Sleep(time.Duration(i+1) * 400 * time.Millisecond)
				case <-kill:
					time.Sleep(time.Duration(i+1) * 10 * time.Millisecond)
					fmt.Print(".")
					return
				}
			}
		}(i, name)
	}

END:
	for {
		select {
		case <-kill:
			fmt.Print("Server shutdown ")
			time.Sleep(time.Duration(len(ws)) * time.Second)
			break END
		}

	}

	fmt.Println(" OK")
}
