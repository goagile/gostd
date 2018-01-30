package main

import (
	"fmt"
	"reflect"
	"math"
)


// Общий геометрический интерфейс
type geomery interface {
	area() float64
	perim() float64
}


// Реализация объекта Прямоугольник
type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return r.width + r.height
}

// Реализация объекта Круг
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}


func print(g geomery) {
	name := reflect.TypeOf(g).Name()
	fmt.Println("geometry =", name)
	fmt.Println("area =", g.area())
	fmt.Println("perim =", g.perim())
}


func main() {
	r := rect{2, 5}
	print(r)

	c := circle{3}
	print(c)
}
