package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var (
	n    = runtime.NumCPU()
	path = "./test/lines.txt"

	lines = make(chan string, n)
	done  = make(chan struct{})

	mu    sync.Mutex
	total = make(map[string]int)
)

func main() {
	fmt.Println("NumCPU: ", n)
	fmt.Println("path: ", path)

	go read()
	for w := 0; w < n; w++ {
		go process()
	}
	wait()

	fmt.Println(total)
}

func read() {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadBytes('\n')
		if err != nil {
			break
		}
		readline(line)
	}
	close(lines)
}

func readline(line []byte) {
	s := string(bytes.TrimSpace(line))
	if s != "" {
		lines <- s
	}
}

func process() {
	for line := range lines {
		for _, f := range strings.Fields(line) {
			count(f)
		}
	}
	done <- struct{}{}
}

func count(f string) {
	mu.Lock()
	defer mu.Unlock()
	total[f]++
}

func wait() {
	for w := 0; w < n; w++ {
		<-done
	}
}
