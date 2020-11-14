package main

import "fmt"

type Numbers chan int

func (ns Numbers) Range(a, b int) {
	defer close(ns)
	for i := a; i < b; i++ {
		ns <- i
	}
}

func (ns Numbers) Print() {
	for n := range ns {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
}

func main() {
	numbers := make(Numbers)

	go numbers.Range(1, 6)

	numbers.Print()
}
