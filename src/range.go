package main

import (
	"fmt"
	"strings"
)


func main() {

	A := [3]int{1, 2, 3}

	fmt.Println("Пропуск идентификатора _, a ")
	
	for _, a := range A {
		fmt.Println(a)
	}



	fmt.Println("Печать идентификатора i, a ")	
	
	for i, a := range A {
		fmt.Println("i = ", i, "a = ", a)
	}


	// обязательно массив бесконечный

	B := []string{"a", "b", "c"}

	fmt.Println("Используем strings.Join")

	fmt.Println(strings.Join(B, " "))

}
