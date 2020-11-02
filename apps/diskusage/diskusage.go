package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

func main() {
	dir := os.Args[1]

	sizes := make(chan int64, 1)
	deadline := time.After(3 * time.Second)
	middleware := time.NewTicker(500 * time.Millisecond)

	var wg sync.WaitGroup
	wg.Add(1)
	go walk(&wg, dir, sizes)
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(sizes)
	}(&wg)

	var nentries int
	var nbytes int64
	done := make(chan struct{})
	go func() {
		for {
			select {
			case s, ok := <-sizes:
				if !ok {
					break
				}
				nentries++
				nbytes += s
			case <-middleware.C:
				log.Printf("%v files %vKB", nentries, nbytes/1e3)
			case <-deadline:
				done <- struct{}{}
			}
		}
	}()
	<-done
}

func walk(wg *sync.WaitGroup, dir string, sizes chan int64) {
	defer wg.Done()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, entry := range entries {
		sizes <- entry.Size()
		if entry.IsDir() {
			wg.Add(1)
			go walk(wg, path.Join(dir, entry.Name()), sizes)
		}
	}
}
