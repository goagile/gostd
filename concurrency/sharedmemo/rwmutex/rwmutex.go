package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.RWMutex
	money int
)

func deposit(amount int) {
	money += amount
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func balance() int {
	return money
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance()
}

func Withdraw(amount int) {
	mu.Lock()
	defer mu.Unlock()
	b := balance()
	if b-amount < 0 {
		fmt.Printf("balance %v amount %v", b, amount)
		return
	}
	money -= amount
}

func main() {
	go Deposit(100)
	go Withdraw(200)

	fmt.Scanln()
	fmt.Println(Balance())
}
