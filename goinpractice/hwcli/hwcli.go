package main

import (
	"fmt"
	"flag"
)

var name = flag.String("name", "World", "A name to say hello to.")
var boost bool

func init() {
	flag.BoolVar(&boost, "boost", false, "A boosted value of string.")
}

func main() {
	flag.Parse()
	if boost == true {
		fmt.Println("............")
		fmt.Printf("Hello, %v\n", *name)
		fmt.Println("............")
	} else {
		fmt.Printf("Hello, %v\n", *name)
	}
}
