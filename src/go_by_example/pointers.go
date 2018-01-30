package main

import "fmt"

func main() {

	x := 1
	fmt.Println("x =", x)
	p := &x
	fmt.Println("&x =", *p)

	incr(&x)
	incr(&x)
	fmt.Println("x incr =", x)

}

func incr(p *int) int {
	*p += 1
	return *p
}
