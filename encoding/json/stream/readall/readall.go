package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../test/example.json")
	if err != nil {
		log.Fatalf("Open: %v\n", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("ReadAll: %v\n", err)
	}
	defer f.Close()
	fmt.Println(string(b))
}
