package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("AAA")
	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatalf("Reader Err: %s\n", err)
	}
	fmt.Printf("line: %q\n", line)
}
