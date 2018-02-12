package main

import "fmt"

func main() {
	// Создаем переменную
	// Под эту переменную выделен адрес в стеке 0xc82000a2e8
	x := new(int)
	// Запишем значение 55 по адресу в стеке 0xc82000a2e8
	*x = 55
	fmt.Println("адрес переменной x:", x)
	fmt.Println("адрес переменной &x:", &x)
	fmt.Println("значение переменной *x:", *x)
}

func newInt() *int {
	return new(int)
}

func newVarInt() *int {
	var dummy int
	return &dummy
}
