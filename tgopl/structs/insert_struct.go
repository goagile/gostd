package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)

	c := Circle{
		Point: Point{X: 1, Y: 2},
		Radius: 5,
	}
	fmt.Printf("%#v\n", c)

	fmt.Printf("X = %v, Y = %v\n", c.Point.X, c.Point.Y)

	// Сокращенный доступ к полям Point
	fmt.Printf("X = %v, Y = %v\n", c.X, c.Y)
}
