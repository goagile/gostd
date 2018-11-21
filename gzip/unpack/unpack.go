package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	for _, file := range os.Args[1:] {
		err := unpack(file)
		if err != nil {
			fmt.Printf("file: %v err: %v\n", file, err)
		}
	}
}

func unpack(gzfile string) error {

	in, err := os.Open(gzfile)
	if err != nil {
		return err
	}
	defer in.Close()

	file := strings.TrimSuffix(gzfile, ".gz")

	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	gzin, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer gzin.Close()

	_, err = io.Copy(out, gzin)
	if err != nil {
		return err
	}

	fmt.Printf("unpacked: %v\n", file)
	return nil
}
