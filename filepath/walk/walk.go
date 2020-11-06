package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	tmp := "tmp"
	tree := "hello/path/test/world"
	path := filepath.Join(tmp, tree)
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatal(err)
	}
	defer func() {
		os.Chdir("../")
		os.RemoveAll(tmp)
	}()
	if err := os.Chdir(tmp); err != nil {
		log.Fatal(err)
	}
	if err := filepath.Walk(".", f); err != nil {
		log.Fatal(err)
	}
}

func f(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}
