package main

import "fmt"

type Person struct {
	name string
}

type Android struct {
	Person
	id string
}

func main() {
	ivan := Person{name: "Ivan"}
	fmt.Println(ivan)

	kasper := Android{id: "QX29172"}
	kasper.name = "Ivan2"
	fmt.Println(kasper)
}
