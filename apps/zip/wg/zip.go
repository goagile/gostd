package main

import (
	"compress/gzip"
	"sync"
	"fmt"
	"os"
	"io"
)

func main() {
	var (
		wg sync.WaitGroup
		file string
		i int
	)
	for i, file = range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			zip(file)
		}(file)
	}
	wg.Wait()
	fmt.Printf("Compressed %v files\n", i)
}

func zip(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Open:%v", err)
	}
	defer f.Close()

	out, err := os.Create(file + ".gzip")
	if err != nil {
		return fmt.Errorf("Create:%v", err)
	}
	defer out.Close()

	w := gzip.NewWriter(out)
	_, err = io.Copy(w, f)
	if err != nil {
		return fmt.Errorf("Copy:%v", err)
	}
	defer w.Close()

	return nil
}
