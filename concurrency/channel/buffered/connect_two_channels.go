package main

import "fmt"

func main() {

	R := []rune{'a', 'b', 'c'}

	A := make(chan rune, 2)
	B := make(chan rune, 2)

	go func() {
		for _, r := range R {
			A<- r
		}
	}()

	go func() {
		for range R {
			B<- <-A
		}
	}()

	for range R {
		r := <-B
		fmt.Printf("%v\n", string(r))
	}	

}
