package main

import "fmt"


func main() {

	counts := make(map[string]int)

	counts["a"] = 1
	counts["b"] = 2

	fmt.Println(counts)
	fmt.Println(counts["a"])
	fmt.Println(counts["b"])

	for k, v := range counts {
		fmt.Printf("%s=%d\n", k, v)
	}
}
