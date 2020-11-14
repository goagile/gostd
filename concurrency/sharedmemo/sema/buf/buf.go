package main

import "fmt"

func main() {
	done := make(chan bool, 1)
	go func() {
		fmt.Println("Hello")
		done <- true
	}()
	<-done
	fmt.Println("Go")
}
