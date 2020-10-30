package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var (
	words = make(chan string)
	wg    sync.WaitGroup
	mp    = make(map[string]int)
	re    = regexp.MustCompile(`[A-Za-z]+`)
	tx    = []string{
		"Hello, Hello, world!",
		"A, Hello, world!!!",
		"Hello, A, world, A A!!!",
	}
)

func main() {
	go func() {
		for {
			select {
			case w := <-words:
				if _, ok := mp[w]; !ok {
					mp[w] = 0
					continue
				}
				mp[w] += 1
			}
		}
	}()
	for _, t := range tx {
		wg.Add(1)
		go func(t string) {
			defer func() {
				wg.Done()
			}()
			r := strings.NewReader(t)
			s := bufio.NewScanner(r)
			for s.Scan() {
				for _, w := range re.FindAllString(s.Text(), -1) {
					words <- w
				}
			}
		}(t)
	}
	wg.Wait()
	fmt.Printf("%v\n", mp)
}
