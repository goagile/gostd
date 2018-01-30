package main

import "fmt"

func main() {

<<<<<<< HEAD
	// пример 1
	fmt.Println()
	x := 1  // объявить переменную x и присвоить ей значение 1
	fmt.Printf("x = %v\n", x)
	p := &x  // адрес переменной x
	fmt.Printf("p = &x = %v\n", *p)  // то, что лежит по адресу p
	*p = 2  // изменить то, что лежит по адресу p 
	fmt.Printf("*p = 2, x = %v\n", x) // напечатать значение переменой х

	// пример 2
	fmt.Println()
	fmt.Printf("getV() == getV(): %v\n", getV() == getV())
	fmt.Println()

	// пример 3
	y := 1;   fmt.Println(y)
	incr(&y); fmt.Println(y)
	incr(&y); fmt.Println(y)
	incr(&y); fmt.Println(y)
}


func getV() *int {
	var v int
	v = 1
	return &v
=======
	x := 1
	fmt.Println("x =", x)
	p := &x
	fmt.Println("&x =", *p)

	incr(&x)
	incr(&x)
	fmt.Println("x incr =", x)

>>>>>>> a1bc0508b218de19424b6e6ca826267080ac74f9
}

func incr(p *int) int {
	*p += 1
	return *p
}
