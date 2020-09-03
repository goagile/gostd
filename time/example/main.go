package main

import (
	"fmt"
	"time"
)

func main() {
	// stamp(time.Now())
	t := time.NewTicker(2 * time.Second)
	for x := 1; x <= 5; x++ {
		<-t.C
		stamp(time.Now())
	}
	t.Stop()
	// stamp(time.Now())
}

func stamp(t time.Time) {
	fmt.Printf("%v:%v\n", t.Minute(), t.Second())
}
