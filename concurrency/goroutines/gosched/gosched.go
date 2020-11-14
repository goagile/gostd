package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("One")
	go func() {
		fmt.Println("Two")
	}()
	fmt.Println("Three")
	runtime.Gosched()
}
