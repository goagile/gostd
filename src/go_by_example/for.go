package main


import "fmt"


func main() {

	fmt.Println("While loop:")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println("For loop:")
	for j := 7; j <= 9; j ++ {
		fmt.Println(j)
	}

	fmt.Println("Бесконечный цикл:")
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println("Вывод только четных")
	for n := 0; n <= 10; n++ {
		if n % 2 == 0 {
			fmt.Println(n)
		}
	}

}
