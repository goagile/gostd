package main

import (
	"fmt"
	"strings"
)


func main() {

	// обязательно массив бесконечный
	B := []string{"a", "b", "c"}
	fmt.Println("Используем strings.Join")
	fmt.Println(strings.Join(B, " "))


	fmt.Println("Используем strings.Split")
	D := strings.Split("a b c d e f g", " ")
	for _, a := range D {
		fmt.Println(a, " ")
	}


}
