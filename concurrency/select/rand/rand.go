package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dice := make([]chan bool, 6)
	for i := 0; i < cap(dice); i++ {
		dice[i] = make(chan bool)
	}

	go func() {
		for {
			dice[rand.Intn(6)] <- true
		}
	}()

	for roll := 0; roll < 36; roll++ {
		var x int
		select {
		case <-dice[0]:
			x = 1
		case <-dice[1]:
			x = 2
		case <-dice[2]:
			x = 3
		case <-dice[3]:
			x = 4
		case <-dice[4]:
			x = 5
		case <-dice[5]:
			x = 6
		}
		fmt.Print(x, " ")
	}
	fmt.Println()

}
