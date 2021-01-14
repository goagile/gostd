package main

import (
	"os"
	"bufio"
	"fmt"
)


func main() {

	var fileName = "test.txt"

	f, err := os.Open(fileName)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "file: %v\n", err)
	}

	printLines(f)

	f.Close()

}


func printLines(f *os.File) {
	
	input := bufio.NewScanner(f)

	for input.Scan() {
		fmt.Println(input.Text())
	}

}
