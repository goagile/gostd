package main

import (
	"fmt"
	"time"
)

func worker(status chan bool, name string) {
	fmt.Println(name, "is start ... ")
	time.Sleep(time.Second * 2)
	fmt.Println(name, "is finish")
	status <- true
}

func main() {

	status := make(chan bool, 4)

	go worker(status, "w1")
	go worker(status, "w2")
	go worker(status, "w3")
	go worker(status, "w4")

	<- status

	time.Sleep(time.Second * 2)

	fmt.Println("<Press any key>")
	input := ""
	fmt.Scanln(&input)

}
