package main

import (
	"fmt"
)

func main() {
	adder := func(a float64) func(float64) float64 {
		s := a
		return func(x float64) float64 {
			s += x
			return s
		}
	}

	addTwo := adder(2)

	fmt.Println(addTwo(4))
}