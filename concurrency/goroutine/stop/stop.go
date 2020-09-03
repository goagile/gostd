package main

import (
	"fmt"
	"time"
)

func main() {

	h := NewHello()

	go h.Run()
	time.Sleep(3*time.Second)
	h.Stop()

	fmt.Println("PAUSE")
	time.Sleep(1*time.Second)	

	go h.Run()
	time.Sleep(5*time.Second)	
	h.Stop()

	fmt.Println("END")

}

func NewHello() *hello {
	return &hello{
		count: 1,
		stop: make(chan bool),
	}
}

type hello struct {
	count int
	stop chan bool
}

func (h *hello) Run() {
	for {
		select {
		default:
			h.Work()
		case <- h.stop:
			return
		}
		h.count++
	}
}

func (h *hello) Work() {
	fmt.Println("HELLO:", h.count)
	time.Sleep(1*time.Second)
}

func (h *hello) Stop() {
	h.stop <- true
}
