package main

import (
	"fmt"
	"sort"
)

func main() {

	ints := []int{7, 2, 1, 6}
	sort.Ints(ints)
	fmt.Println("ints:", ints)

	strs := []string{"z", "b", "f", "a"}
	sort.Strings(strs)
	fmt.Println("strs:", strs)

}
