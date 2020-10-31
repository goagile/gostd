package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 5
	s := NewStore()

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			k := fmt.Sprintf("%v_%v", "k", i)
			s.Set(k, i)
		}(i)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		k := fmt.Sprintf("%v_%v", "k", i)
		v := s.Get(k)
		fmt.Println(k, "->", v)
	}
}

type store map[string]interface{}

func NewStore() store {
	s := make(store)
	return s
}

func (s store) Set(k string, v interface{}) {
	s[k] = v
}

func (s store) Get(k string) interface{} {
	return s[k]
}
