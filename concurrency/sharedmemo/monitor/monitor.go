package main

import (
	"fmt"
	"time"
)

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func monitor() {
	var deposit int
	for {
		select {
		case amount := <-deposits:
			deposit += amount
		case balances <- deposit:
		}
	}
}

func main() {

	go monitor()

	fmt.Println("AFTER", Balance())

	go func() {
		Deposit(100)
		fmt.Println("Deposit 100")
	}()

	go func() {
		Deposit(200)
		fmt.Println("Deposit 200")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("BEFORE", Balance())
}
