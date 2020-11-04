package main

import "fmt"

var (
	sema    = make(chan struct{}, 1) // сторожит balance
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // захватывает маркер
	balance = balance + amount
	<-sema // отпускает маркер
}

func main() {
	go Deposit(100)
	go Deposit(200)

	fmt.Print("press enter ")
	fmt.Scanln()
	sema <- struct{}{} // захватывает маркер
	b := balance
	<-sema // освобождает маркер
	fmt.Println("balance = ", b)
}
