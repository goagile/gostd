package main

import (
	"fmt"
)

func main() {
	

	var i interface{}

	i = "Hello"

	var s string
	s = i.(string)
	fmt.Println(s)

	var ok bool

	var z string
	z, ok = i.(string)
	fmt.Println(z, ok)	

	var f float64
	f, ok = i.(float64)
	fmt.Println(f, ok)		

	f = i.(float64)
	fmt.Println(f)

}