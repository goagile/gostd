package main

import (
	"errors"
	"fmt"
	"time"
)

type num struct {
	val int
	err error
}

func main() {

	sqrs := make(chan num)

	n := 20

	for i := 0; i < n; i++ {
		go func(i int) {
			var x num
			x.val, x.err = square(i)
			sqrs <- x
		}(i)
	}

	for i := 0; i < n; i++ {
		x := <-sqrs
		if x.err != nil {
			fmt.Println(x.err)
			break
		}
		fmt.Println(x.val)
	}
}

func square(i int) (int, error) {
	x := i * i
	time.Sleep(time.Duration(x) * time.Millisecond)
	if x == 25 {
		return 0, errors.New("ERROR WITH: x = 25")
	}
	return x, nil
}
