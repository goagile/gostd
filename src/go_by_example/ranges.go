package main

import "fmt"

func main() {
	
	// создаем массив с помощью литерала
	A := []int{3, 2, 1}
	
	fmt.Println("Пройдем по массиву в цикле")
	for i := 0; i < len(A); i++ {
		fmt.Printf("A[%v] = %v\n", i, A[i])
	}

	fmt.Println("Пройдем по массиву в цикле с помощью range")
	for i, a := range A {
		fmt.Printf("A[%v] = %v\n", i, a)
	}

	// Создаем Hash
	fmt.Println("Создаем hash")
	H := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(H)

	// Бегаем по хэшу
	fmt.Println("Бегаем по хэшу")
	for k, v := range H {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	// Бегаем по строке
	fmt.Println("Бегаем по строке")
	S := "golang"
	for i, k := range S {
		fmt.Printf("i = %v, k = %v\n", i, k)
	}

}
