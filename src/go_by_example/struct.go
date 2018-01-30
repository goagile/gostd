package main

import "fmt"

type person struct {
	name string
	salary int
}

func main() {

	p := person{name: "Bob", salary: 32}
	fmt.Println(p)

	incr(&p)
	fmt.Println(p)

}

func incr(p *person) {
	p.salary += 100
}
