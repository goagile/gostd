package main

import (
	"sort"
	"fmt"
)

func main() {

	n := []int{2, 4, 5, 9, 1, 3, 6}
	
	sort.Ints(n)

	fmt.Println(n)

}
