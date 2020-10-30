package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	n := 8
	lock := make(chan struct{}, 1)
	for i := 0; i < n; i++ {
		go func(i int) {
			fmt.Println(tab(i), "(T)", tab(n-i))
			lock <- struct{}{}
			fmt.Println(tab(i), "(L)", tab(n-i))
			time.Sleep(time.Duration(i*100) * time.Millisecond)
			<-lock
			fmt.Println(tab(i), "(U)", tab(n-i))
		}(i)
	}
	time.Sleep(3 * time.Second)
}

func tab(n int) string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, "  |  ")
	}
	return strings.Join(s, "")
}
