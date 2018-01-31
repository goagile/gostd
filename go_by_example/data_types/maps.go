package main

import "fmt"

func main() {

	// Создание хэша с помощью функции make
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("m = ", m)

	v1 := m["k1"]
	fmt.Println("v1 =", v1)
	fmt.Println("len m =", len(m))

	delete(m, "k2")
	fmt.Println("m =", m)

	_, is_exist := m["k1"]
	fmt.Println(is_exist)

	// Создание хэша без функции make
	n := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("n =", n)

}
