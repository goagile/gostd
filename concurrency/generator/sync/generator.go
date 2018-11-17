package main

import "fmt"

func main() {
	for _, i := range genInts(10) {
		fmt.Println(i)
	}
}

func genInts(n int) []int {
	result := []int{}
	for i := 0; i < 10; i++ {
		result = append(result, i)
	}
	return result
}
