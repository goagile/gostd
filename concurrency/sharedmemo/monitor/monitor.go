package main

import (
	"fmt"
	"time"
)

var (
	deposits  = make(chan int)
	balances  = make(chan int)
	withdraws = make(chan int)
	zero      = make(chan Status)
)

type Status struct {
	Balance, Amount int
}

func Withdraw(amount int) {
	balance := Balance()
	if balance-amount < 0 {
		zero <- Status{balance, amount}
		return
	}
	withdraws <- amount
}

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func accountmonitor() {
	var deposit int
	for {
		select {
		case z := <-zero:
			fmt.Printf("Не хватает средств для снятия Amount %v Balance %v\n", z.Amount, z.Balance)
		case amount := <-withdraws:
			deposit -= amount
		case amount := <-deposits:
			deposit += amount
		case balances <- deposit:
		}
	}
}

func main() {
	go accountmonitor()
	fmt.Println("Balance", Balance())

	go Deposit(1000)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Положили 100, Balance", Balance())

	go Deposit(200)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Положили 200, Balance", Balance())

	go Withdraw(500)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Сняли 500, Balance", Balance())

	go Withdraw(2000)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Balance", Balance())
}
