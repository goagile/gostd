package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // защищает balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	go Deposit(100)
	go Deposit(200)
	go Deposit(300)

	fmt.Print("press enter")
	fmt.Scanln()
	fmt.Println(Balance())
}
