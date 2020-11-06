package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type fileinfos chan fileinfo

type fileinfo struct {
	path string
	hash []byte
}

func main() {
	infos := make(fileinfos)
	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatal("need dir")
	}
	dir := os.Args[1]

	go findDups(dir, infos)

	result := merge(infos)

	for _, v := range result {
		fmt.Println(v.filename)
		for _, x := range v.paths {
			fmt.Println("\t", x)
		}
	}
}

type pathdata struct {
	filename string
	paths    []string
}

func merge(infos fileinfos) map[string]*pathdata {
	m := make(map[string]*pathdata)
	for i := range infos {
		key := fmt.Sprintf("%x", i.hash)
		if _, ok := m[key]; ok {
			m[key].paths = append(m[key].paths, i.path)
			continue
		}
		m[key] = &pathdata{
			filename: i.path,
			paths:    make([]string, 0),
		}
	}
	return m
}

func findDups(dir string, infos fileinfos) {
	defer close(infos)
	var wg sync.WaitGroup
	w := makeWalkFunc(infos, &wg)
	if err := filepath.Walk(dir, w); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

func makeWalkFunc(infos fileinfos, wg *sync.WaitGroup) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				process(infos, path, info)
			}(wg)
		}
		return err
	}
}

func process(infos fileinfos, path string, info os.FileInfo) {
	f, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}

	hash := sha1.New()
	if _, err := io.Copy(hash, f); err != nil {
		log.Println(err)
		return
	}

	infos <- fileinfo{
		path: path,
		hash: hash.Sum(nil),
	}
}
