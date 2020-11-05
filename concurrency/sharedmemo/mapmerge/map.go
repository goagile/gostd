package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

const path = "./test/lines.txt"

var (
	n       = runtime.NumCPU()
	lines   = make(chan string, n)
	results = make(chan map[string]int)
	done    = make(chan struct{}, n)
	total   = make(chan map[string]int)
)

func main() {
	fmt.Println("NumCPU: ", n)
	fmt.Println("Read", path)

	go read()
	for w := 0; w < n; w++ {
		go process()
	}
	go wait()
	go merge()

	fmt.Println(<-total)
}

func merge() {
	t := make(map[string]int)
	for r := range results {
		for k, v := range r {
			t[k] += v
		}
	}
	total <- t
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
		readline(lines, line)
	}

	close(lines)
}

func readline(lines chan<- string, line []byte) {
	s := string(bytes.TrimSpace(line))
	if s != "" {
		lines <- s
	}
}

func process() {
	for line := range lines {
		processline(results, line)
	}
	done <- struct{}{}
}

func processline(results chan map[string]int, line string) {
	m := make(map[string]int)
	for _, f := range strings.Fields(line) {
		m[f]++
	}
	results <- m
}

func wait() {
	for w := 0; w < n; w++ {
		<-done
	}
	close(results)
}
