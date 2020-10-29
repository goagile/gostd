package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"io"
)

func main() {
	args := os.Args[1:]

	files := make(chan string)
	go func() {
		for _, f := range args {
			files<- f
		}
		close(files)
	}()

	n := len(args)

	done := make(chan struct{}, n)
	
	for i := 0; i < n; i++ {
		go func() {
			err := zip(<-files)
			if err != nil {
				fmt.Println("zip", err)
			}
			done<- struct{}{}
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

	fmt.Printf("Compressed %v files\n", n)
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
