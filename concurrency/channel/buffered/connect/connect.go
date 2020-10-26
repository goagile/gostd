package main

import "fmt"

func main() {

	runes := []rune{'a', 'b', 'c'}

	in  := make(chan rune, 3)
	out := make(chan rune, 3)

	go func() {
		for _, r := range runes {
			in <- r
		}
		close(in)
	}()

	go func() {
		for x := range in {
			out <- x
		}
		close(out)
	}()

	for o := range out {
		fmt.Printf("%v\n", string(o))
	}	

}
