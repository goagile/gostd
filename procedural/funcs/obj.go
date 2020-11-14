package main

import (
	"fmt"
	"math"
)


type F func(float64, float64) float64


func main() {
	
	hypot := func(a, b float64) float64 {
		return math.Sqrt(a*a + b*b)
	}

	compute := func(fn F) float64 {
		return fn(2, 3)
	}

	fns := []F{
		hypot,
		math.Pow,
	}

	for _, fn := range fns {
		fmt.Println(compute(fn))
	}


}
