package main

import "fmt"

func main() {

	// Пустой массив
	var a [5]int
	fmt.Printf("empty array: %v\n", a)

	// Получение и запись
	a[4] = 100
	fmt.Printf("set: %v\n", a)
	fmt.Printf("get: %v\n", a[4])

	// Длина массива
	fmt.Printf("len A: %v\n", len(a))

	// Литеральное начальное заполнение
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("b = %v\n", b)

	// Двумерный массив
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
