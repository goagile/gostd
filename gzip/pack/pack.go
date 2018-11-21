package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	for _, file := range os.Args[1:] {
		err := pack(file)
		if err != nil {
			fmt.Printf("file: %v err: %v\n", file, err)
		}
	}
}

func pack(file string) error {

	in, err := os.Open(file)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(file + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gzout := gzip.NewWriter(out)
	_, err = io.Copy(gzout, in)
	if err != nil {
		return err
	}
	defer gzout.Close()

	fmt.Printf("compressed: %v\n", out.Name())
	return nil
}
