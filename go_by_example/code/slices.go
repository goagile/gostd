package main

import "fmt"

func main() {

	// Срезы создаются с помощью make
	s := make([]string, 3)

	// У среза есть длина
	fmt.Println("len s = ", len(s))

	// В срез можно добавлять элементы по индексу
	// Индекс не может быть больше длины среза
	s[0] = "A"
	s[1] = "B"
	fmt.Println("s = ", s)
	fmt.Println("s[0] = ", s[0])

	// Добавить в конец среза
	s = append(s, "x")
	s = append(s, "xa", "xax")
	fmt.Println("s = ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy of s is c = ", c)

	fmt.Println("\nСрезы:")
	fmt.Println("c =", c)
	fmt.Println("c[1:] =", c[1:])
	fmt.Println("c[2:] =", c[2:])
	fmt.Println("c[3:] =", c[3:])
	fmt.Println("c[4:] =", c[4:])
	fmt.Println("c[5:] =", c[5:])

	fmt.Println("\nСрезы обра с :")
	fmt.Println("c[5:] =", c[:1])

	

}
