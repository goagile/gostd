package main

import "fmt"

var (
	ints = []int{1, 2, 3}
	msgs = make(chan int)
	done = make(chan bool)
)

func main() {
	go produce()
	go consume()
	<-done
}

func produce() {
	for _, i := range ints {
		msgs <- i
	}
	done <- true
}

func consume() {
	for {
		fmt.Println(<-msgs)
	}
}
