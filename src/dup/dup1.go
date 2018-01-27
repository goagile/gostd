package main

import (
	"fmt"
	"bufio"
	"os"
)


func main() {

	counter := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counter[input.Text()]++
	}

	for k, v := range counter {
		fmt.Printf("%s = %d\n", k, v)	
	}

}
