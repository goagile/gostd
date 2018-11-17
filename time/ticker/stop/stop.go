package main

import (
	"fmt"
	"time"
)

func main() {

	// set ticker to tick every second
	ticker := time.NewTicker(time.Second)

	// run ticker inn separate goroutine
	go func() {
		for tick := range ticker.C {
			fmt.Printf("🕐 %v:%v:%v\n", tick.Hour(), tick.Minute(), tick.Second())
		}
	}()

	// stop the ticker after 2 sec
	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
