package main

import (
	"fmt"
	"time"
)

const N = 5

var (
	sm = NewSafeMap()
)

func main() {

	go func() {
		sm.Put("A", 555)
	}()

	go func() {
		sm.Put("B", 234)
	}()

	go func() {
		sm.Update("B", func(v interface{}, ok bool) interface{} {
			if ok {
				return 100
			}
			return 0
		})
	}()

	time.Sleep(500 * time.Millisecond)

	v, _ := sm.Get("A")
	fmt.Println("A", v)

	v, _ = sm.Get("B")
	fmt.Println("B", v)

	sm.Update("B", func(v interface{}, ok bool) interface{} {
		if ok {
			return v.(int) + 1000
		}
		return 0
	})

	v, _ = sm.Get("B")
	fmt.Println("B", v)

	v, _ = sm.Get("C")
	fmt.Println("C", v)
}

type safemap struct {
	put, get, update chan item
}

type item struct {
	k  string
	v  interface{}
	ok bool
	r  chan item
	u  func(v interface{}, ok bool) interface{}
}

func NewSafeMap() *safemap {
	sm := &safemap{
		put:    make(chan item),
		get:    make(chan item),
		update: make(chan item),
	}
	go sm.monitor()
	return sm
}

func (sm *safemap) monitor() {
	m := make(map[string]interface{})

	for {
		select {

		case i := <-sm.put:
			m[i.k] = i.v

		case i := <-sm.get:
			v, ok := m[i.k]
			i.r <- item{v: v, ok: ok}

		case i := <-sm.update:
			v, ok := m[i.k]
			m[i.k] = i.u(v, ok)
		}
	}
}

func (sm *safemap) Put(k string, v interface{}) {
	sm.put <- item{k: k, v: v}
}

func (sm *safemap) Get(k string) (interface{}, bool) {
	r := make(chan item)
	sm.get <- item{k: k, r: r}
	i := <-r
	return i.v, i.ok
}

func (sm *safemap) Update(k string, u func(v interface{}, ok bool) interface{}) {
	sm.update <- item{k: k, u: u}
}
