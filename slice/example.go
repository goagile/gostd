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


	var z []int

	fmt.Println("z == nil is ", z == nil) 
	fmt.Println("z len, cap = ", len(z), cap(z))


	b := make([]int, 0, 3)
	fmt.Println("b len, cap = ", len(b), cap(b))

	a1 := []int{0,1,2,3,4}
	for i, v := range a1 {
		fmt.Println("i:", i, " v:", v)
	}

}
