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

type store struct {
	mu   sync.Mutex             // <- mu защищает data
	data map[string]interface{} // <- data защищается mu
}

func NewStore() *store {
	return &store{
		data: make(map[string]interface{}),
	}
}

func (s *store) set(k string, v interface{}) {
	s.data[k] = v
}

func (s *store) Set(k string, v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.set(k, v)
}

func (s *store) get(k string) interface{} {
	return s.data[k]
}

func (s *store) Get(k string) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.get(k)
	return v
}
