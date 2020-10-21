package main

import (
	"fmt"
	"sort"
)

func main() {

	i := []int{4, 3, 2, 1, 0}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"z", "c", "a", "x", "b"}
	sort.Strings(s)
	fmt.Println(s)

	f := []float64{0.56, 0.45, 0.33, 0.1}
	sort.Float64s(f)
	fmt.Println(f)

	r := []int{2, 3, 4, 5}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	fmt.Println(r)
}
