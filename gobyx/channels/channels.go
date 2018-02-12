package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pong"
	}()

	go func() {
		messages <- "pang"
	}()

	fmt.Println(<- messages)
	fmt.Println(<- messages)
	fmt.Println(<- messages)

}
