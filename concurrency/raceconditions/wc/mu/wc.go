package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

var (
	wg sync.WaitGroup
	mu sync.Mutex
	mp = make(map[string]int)
	re = regexp.MustCompile(`[A-Za-z]+`)
	tx = []string{
		"Hello, Hello, world!",
		"A, Hello, world!!!",
		"Hello, A, world, A A!!!",
	}
)

func main() {
	for _, t := range tx {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			r := strings.NewReader(t)
			s := bufio.NewScanner(r)
			for s.Scan() {
				for _, w := range re.FindAllString(s.Text(), -1) {
					mu.Lock()
					_, ok := mp[w]
					mu.Unlock()
					if !ok {
						mu.Lock()
						mp[w] = 1
						mu.Unlock()
						continue
					}
					mu.Lock()
					mp[w] += 1
					mu.Unlock()
				}
			}
		}(t)
	}
	wg.Wait()
	fmt.Printf("%v\n", mp)
}
