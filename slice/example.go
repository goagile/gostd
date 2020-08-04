package main


import (
	"fmt"
)


func main() {


	s := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Println("s[:] = ", s[:])
	fmt.Printf("s type is %T\n", s)
	fmt.Println("s len, cap = ", len(s), cap(s))


	a := s[:3]

	fmt.Printf("a type is %T\n", a)
	fmt.Println("a = s[:3] len, cap = ", len(a), cap(a))


}

