package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(134)
	for i := 0; i < 3; i++ {
		fmt.Println(rand.Int())
	}
}
