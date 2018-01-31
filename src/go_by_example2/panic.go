package main

import (
	// "fmt"
	"os"
)

func main() {

	// panic("a problem")

	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}

}
