package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan int)

	go worker(1, done)

	<- done

}

func worker(n int, done chan<- int) {
	fmt.Println(n, "working ... ")
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(n, "done")
	done <- n
}
