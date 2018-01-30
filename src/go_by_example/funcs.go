package main

import "fmt"


func plus(a, b int) int {
	return a + b
}

func multy() (int, int) {
	return 2, 3
}

func sum(nums ...int) int {
	result := 0
	for _, a := range nums {
		result += a
	}
	return result
}

func main() {

	res := plus(1, 2)
	fmt.Println("res =", res)

	a, b := multy()
	fmt.Printf("a = %v, b = %v\n", a, b)

	_, d := multy()
	fmt.Printf("d = %v\n", d)

    nums := []int{1, 2, 3, 4}
    res = sum(nums...)
	fmt.Printf("res = %v\n", res)

}
