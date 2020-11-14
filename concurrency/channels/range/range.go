package main

import "fmt"

func main() {

	Q := make(chan string)

	go func() {
		Q<- "one"
		Q<- "two"
		Q<- "three"
		close(Q)
	}()

	for e := range Q {
		fmt.Println(e)
	}

}
