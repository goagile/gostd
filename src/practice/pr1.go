package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)


func main() {

	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}


	lines := ""
	input := bufio.NewScanner(f)
	for input.Scan() {
		lines += input.Text()
	}
	splitted := strings.Split(string(lines), ";")


	persons := make(map[string]string)
	for _, line := range splitted {
		if string(line) == "" {
			continue
		}
		kv := strings.Split(string(line), " = ")
		k := kv[0]
		v := kv[1]
		persons[k] = v
	}
	fmt.Println(persons)

}
