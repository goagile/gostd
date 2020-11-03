package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	m = 5
	n = 4
)

func main() {
	fmt.Println()
	for i := 0; i < n; i++ {
		go do(i)
	}
	fmt.Scanln()
}

func do(i int) {
	for j := 1; j <= m; j++ {
		fmt.Println(strings.Repeat("\t", i), i, strings.Repeat("*", j))
		time.Sleep(time.Duration(j) * 100 * time.Millisecond)
		runtime.Gosched()
	}
}
