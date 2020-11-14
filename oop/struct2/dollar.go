package main

import (
	"fmt"
)


type Dollar float64

func (d Dollar) String() string {
	return fmt.Sprintf("$%v", float64(d))
}

func main() {
	five := Dollar(5)
	fmt.Println(five + five)
}
