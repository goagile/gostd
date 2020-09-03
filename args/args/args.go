package main 

import (
	"os"
	"fmt"
)

func main() {

	fmt.Println("Args")

	for _, a := range os.Args {
		fmt.Println(a)
	}

}