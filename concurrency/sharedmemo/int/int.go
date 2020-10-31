package main

import (
	"fmt"
	"sync"
)

func main() {
	go share()

	set(123)
	set(123)
	fmt.Println(get())
	fmt.Println(get())
	fmt.Println(get())

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		set(534)
		fmt.Println("\t", get())
		fmt.Println("\t", get())
	}()
	wg.Wait()

	set(45638)
	set(146)
	fmt.Println(get())
}

//
// share monitor
//
func share() {
	var v int
	for {
		select {
		case v = <-in:
		case out <- v:
		}
	}
}

//
// set
//
var in = make(chan int)

func set(v int) {
	in <- v
}

//
// get
//
var out = make(chan int)

func get() int {
	return <-out
}
