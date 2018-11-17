package main

import (
	"fmt"
	"time"
)

func main() {

	// set ticker to tick every second
	ticker := time.NewTicker(time.Second)

	for tick := range ticker.C {
		fmt.Printf("🕐 %v:%v:%v\n", tick.Hour(), tick.Minute(), tick.Second())
	}

	ticker.Stop()

}
