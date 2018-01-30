package main

import (
	"fmt"
	"errors"
)


func process(a int) (int, error) {
	if a < 0 {
		return -1, errors.New("Число меньше 0")
	}
	return a + 1, nil
}


func main() {
	seq := []int{1, -2, 3, 6, 6, -4, 1, 2, -4}
	for _, a := range seq {
		x, err := process(a)
		if err != nil {
			fmt.Println(x, err)
		} else {
			fmt.Println(x)
		}
	}
}
