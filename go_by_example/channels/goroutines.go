package main

import "fmt"


func do(n []int) {
	for _, a := range n {
		fmt.Println(a)
	}
}

func boo(symbol string) int {
	s := 0
	for a := 0; a < 10; a++ {
		s += a
		fmt.Println(symbol, s)
	}
	return s
}

func main() {
	// go do([]int{1, 2, 3})	
	// go do([]int{4, 5, 6, 7})
	// go do([]int{10, 20, 30})

	go boo("a")
	go boo("b")
	go boo("c")

	input := ""
	fmt.Scanln(&input)
}
