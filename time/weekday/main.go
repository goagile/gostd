package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Weekday()
	fmt.Println(t)
}
