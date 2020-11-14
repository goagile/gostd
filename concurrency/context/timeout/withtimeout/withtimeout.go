package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	go work(ctx)

	fmt.Scanln()

	cancel()

	fmt.Println("OK")

}

func work(ctx context.Context) {
	for {
		select {
		default:
			fmt.Print("A ")
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			fmt.Print("Done by timeout")
			return
		}
	}
}
