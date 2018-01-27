package main

import (
	"os"
	"bufio"
	"fmt"
)


func main() {

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		fmt.Println(input.Text())
	}

}
