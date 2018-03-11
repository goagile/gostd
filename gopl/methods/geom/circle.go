package geom

import "math"

type Circle struct {
	X, Y, R float64
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.R
}
