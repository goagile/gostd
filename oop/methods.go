package main

import "fmt"

type person struct {
	name string
	salary int
}

func (self *person) incr() {
	self.salary += 100
}

func (self *person) getName() string {
	return "|" + self.name + "|"
}

func main() {

	p := person{name: "Bob", salary: 200}
	p.incr()
	fmt.Println(p)
	fmt.Println(p.getName())
	
}
