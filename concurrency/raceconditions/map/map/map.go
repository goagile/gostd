package main

import (
	"fmt"
	"time"
)

//
// go run --race race.go
//
func main() {
	m := map[string]string{}

	go func() {
		m["message"] = "Hello 1"
	}()

	go func() {
		m["message"] = "Hello 2"
	}()

	time.Sleep(time.Second)

	fmt.Println(m["message"])
}
