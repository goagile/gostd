package main

import "fmt"

func main() {

	channel := make(chan string, 2)

	channel <- "buffered"
	channel <- "channel"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	// fmt.Println(<- channel) fatal error: all goroutines are asleep - deadlock!

}
