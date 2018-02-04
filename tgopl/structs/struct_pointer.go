package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	
	// (1) 
	p := Point{1, 2}
	fmt.Println("p:", p)

	// (2)
	pp := &Point{1, 2}
	fmt.Println("pp:", pp)

	// (3)
	ppp := new(Point)
	*ppp = Point{1, 2}
	fmt.Println("ppp:", ppp)

	// (4)
	p2 := move(p)
	fmt.Println("копия передается в функцию move:", p2)
	fmt.Println("исходная переменная", p)

	p3 := movePointer(&p)
	fmt.Println("ссылка передается в функцию movePointer:", p3)	
	fmt.Println("исходная переменная", p)

}

func move(p Point) Point {
	p.X += 1
	p.Y += 1
	return p
}

func movePointer(p *Point) *Point {
	p.X += 1
	p.Y += 1
	return p
}
