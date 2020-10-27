package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	tree(os.Stdout, os.Args[1])
}

func tree(w io.Writer, dir string) {
	fmt.Fprintln(w, dir)
	walk(w, dir, "")
}

func walk(w io.Writer, dir string, tab string) {
	fs, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	last := len(fs) - 1
	for i, f := range fs {
		fmt.Fprintf(w, "%v%v\n", tab+t(i == last), f.Name())
		if f.IsDir() {
			walk(w, path.Join(dir, f.Name()), tab+d(i == last))
		}
	}
}

func t(last bool) string {
	if last {
		return "└── "
	}
	return "├── "
}

func d(last bool) string {
	if last {
		return "    "
	}
	return "│   "
}
