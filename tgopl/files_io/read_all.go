package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)


func main() {

	filename := "test.txt"

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	lines := strings.Split(string(data), "\n")

	for _, a := range lines {
		fmt.Println(a)
	}

}
