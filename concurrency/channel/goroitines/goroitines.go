package main

import "fmt"


func main() {

	//
	// run concurrent
	//
	go routine("A")
	go routine("\tB")
	go routine("\t\tC")

	// 
	// wait for user
	//
	fmt.Scanln()

}

func routine(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
	}
}
