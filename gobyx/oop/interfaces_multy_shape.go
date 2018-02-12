package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	str() string
}

type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) str() string {
	return fmt.Sprintf("Circle{r: %v}", c.r)
}

type MultyShape struct {
	shapes []Shape
}

func (m *MultyShape) area() float64 {
	var result float64
	result = 0
	for _, s := range m.shapes {
		result += s.area()
	}
	return result
}

func (m *MultyShape) str() string {
	return fmt.Sprintf("MultyShape{len=%v}", len(m.shapes))
}

func (m *MultyShape) append(s Shape) {
	m.shapes = append(m.shapes, s)
}

func main() {
	c := Circle{r: 5}
	fmt.Println(c.str())
	fmt.Println("c.area() =", c.area())

	m := MultyShape{}
	m.append(&c)
	m.append(&c)
	fmt.Println(m.str())
	fmt.Println("m.area() =", m.area())	
}
