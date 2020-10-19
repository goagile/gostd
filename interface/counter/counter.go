package main

import "fmt"

func main() {
	var c ByteCounter = 0

	n, _ := c.Write([]byte("hola"))
	fmt.Println("len('hola' =", n)

	c = 0

	fmt.Fprintf(&c, "%v", "hello")
	fmt.Println("len(hello) =", c)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return int(*c), nil
}
