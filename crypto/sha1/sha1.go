package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strings"
)

func main() {
	m := make(map[string]string)
	for _, value := range []string{
		`{"name":"A", "count":1}`,
		"AAAAA",
		"BBBBBBB",
		"Hello, world!",
		`{"name":"A", "count":1}`,
	} {
		hash := sha1.New()
		io.Copy(hash, strings.NewReader(value))
		key := fmt.Sprintf("%X", hash.Sum(nil))
		if _, ok := m[key]; ok {
			fmt.Println("(!!!) DUPLICATE: ", key, value)
			continue
		}
		m[key] = value
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
