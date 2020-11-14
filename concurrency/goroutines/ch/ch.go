package main

import "fmt"

func main() {

	ch := make(chan string)

	go func(ch chan string) {
		ch <- "Hello 1"
	}(ch)

	go func(ch chan string) {
		ch <- "Hello 2"
	}(ch)

	go func(ch chan string) {
		ch <- "Hello 3"
	}(ch)

	for i := 1; i <= 3; i++ {
		x := <-ch
		fmt.Println(x)
	}

}
