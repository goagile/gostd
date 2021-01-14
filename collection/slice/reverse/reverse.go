package main

import (
	"fmt"
)

func main() {
	for _, a := range []Ints{
		{},
		{1, 2},
		{1, 2, 3},
	} {
		a.Reverse()
		fmt.Println(a)
	}
}

type Ints []int

func (a Ints) Reverse() {
	for i := 0; i <= len(a)/2-1; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
}
