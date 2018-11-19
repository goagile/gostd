package main

import (
	"flag"
	"fmt"
)

var name = flag.String("n", "World", "A name to say hello to.")

func main() {
	flag.Parse()

	fmt.Printf("Hello, %v\n", *name)
}
