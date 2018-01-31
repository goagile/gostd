package main

import "fmt"


type Person struct {
	name string
	age int
}


func main() {

	// fmt.Println(Person{"Bob", 20})

	// fmt.Println(Person{name: "Antony", age: 32})

	// fmt.Println(Person{name: "Sergey"})

	// fmt.Println(&Person{name: "Pointer", age: 12})

	s := Person{name: "Ref", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)
	fmt.Println(sp.age)

	z := s
	s.age = 52
	fmt.Println(z.age)
	fmt.Println(s.age)
	fmt.Println(sp.age)

}
