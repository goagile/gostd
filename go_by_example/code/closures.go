package main

import "fmt"

func genNext() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {


	next := genNext()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	
}
