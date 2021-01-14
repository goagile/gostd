package main

import (
	"fmt"

	"github.com/goagile/gostd/oop/methods/geom"
)

func main() {
	p := geom.Point{10, 20}
	q := geom.Point{30, 40}
	distance := p.Distance(q)
	fmt.Printf("Point.Distance = %v\n", distance)

	c := new(geom.Circle)
	c.X = 10
	c.Y = 10
	c.R = 20
	perim := c.Perim()
	fmt.Printf("Circle.Perim = %v\n", perim)
}
