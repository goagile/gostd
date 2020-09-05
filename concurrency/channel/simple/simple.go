package main

import "fmt"

func main() {

	tube := make(chan string)

	go func() {
		tube<- "a"
	}()

	a := <-tube

	fmt.Println(a)
	
}
