package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go work(ctx)

	fmt.Scanln()

	cancel()

	time.Sleep(1 * time.Second)
}

func work(ctx context.Context) {
	for {
		select {
		default:
			fmt.Print("work ")
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			fmt.Println("work Done")
			return
		}
	}
}
