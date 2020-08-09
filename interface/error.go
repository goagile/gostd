package main

import (
	"fmt"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	return &MyErr{
		What: "i did not work",
		When: time.Now(),
	}
}

type MyErr struct {
	What string
	When time.Time
}

func (err* MyErr) Error() string {
	return fmt.Sprintf("at %v: %v", err.When, err.What)
}
