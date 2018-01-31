package main

import "fmt"

func main() {

	pings := make(chan string, 2)
	pongs := make(chan string, 2)

	pings<- "a"
	pings<- "b"

	pongs<- <-pings
	pongs<- <-pings

	fmt.Println(<-pongs)
	fmt.Println(<-pongs)

}
