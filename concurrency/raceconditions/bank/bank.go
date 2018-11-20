package main

import (
	"fmt"
	"sync"
)

var (
	balance int
)

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	var wg sync.WaitGroup

	// Bob
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Deposit(1)
			fmt.Printf("Bob balance = %v\n", Balance())
		}()
	}

	// Alice
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(1000)
		fmt.Printf("Alice balance = %v\n", Balance())
	}()

	wg.Wait()
}
