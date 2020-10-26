package main

import (
	"fmt"
)

type Water int

var L Water = 1
	
func main() {

	tube := make(chan Water, 3)


	go func() {
		tube<-  5*L
		tube<- 15*L
		tube<- 10*L
		close(tube)
	}()


	var barrel Water
	barrel += <-tube
	barrel += <-tube


	var bootle Water
	bootle += <-tube


	fmt.Println("barrel", barrel)
	fmt.Println("bottle", bootle)
	
}
