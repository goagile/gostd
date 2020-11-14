package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func (p Point) ShiftX(dx float64) {
	p.X += dx
}	

func (p* Point) ShiftXPointer(dx float64) {
	p.X += dx
}	

func abs(p Point) float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func shiftx(p Point, dx float64) {
	p.X += dx
}

func main() {
	p := Point{4, 5}
	fmt.Println(p.Abs())
	p.ShiftX(10)
	fmt.Println(p.Abs())
	p.ShiftXPointer(10)
	fmt.Println(p.Abs())

	b := Point{4, 5}
	fmt.Println(abs(b))
	shiftx(b, 10)
	fmt.Println(abs(b))
}
