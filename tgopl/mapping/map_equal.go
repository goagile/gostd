package main

import "fmt"

func main() {

	g1 := map[string]int{
		"socks": 30,
		"shoes": 10, 
	}

	g2 := map[string]int{
		"socks": 30,
		"shoes": 10, 
	}

	g3 := map[string]int{}

	g4 := map[string]int{
		"socks": 40,
		"shoes": 20,
	}

	fmt.Println("g1 equal g2:", equal(g1, g2))
	fmt.Println("g1 equal g3:", equal(g1, g3))
	fmt.Println("g1 equal g4:", equal(g1, g4))

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
