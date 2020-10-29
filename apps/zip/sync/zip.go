package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"io"
)

func main() {
	files := os.Args[1:]
	var n int
	for _, file := range files {
		err := zip(file)
		if err != nil {
			fmt.Println("zip", err)
			continue
		}
		n++
	}
	fmt.Println("Compressed %v files", n)
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
