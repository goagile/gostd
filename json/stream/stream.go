package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	"log"
	"encoding/json"
)

func main() {

	s, err := NewPersonScanner("./test/example.json")
	if err != nil {
		log.Fatalf("NewPersonScanner %v\n", err)
	}

	for s.Next() {
		p, err := r.Scan()
		if err != nil {
			log.Fatalf("Scan err: %v\n", err)
		}
		fmt.Println(p)
	}

	fmt.Printf("Read %v persons.\n", r.Count())

}

func NewPersonScanner(filename string) (*PersonScanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := &PersonScanner{
		dec: json.NewDecoder(f),
		next: true,
		scrolled: false,
	}
	return r, nil
}

type PersonScanner struct {
	dec *json.Decoder
	count int
	file *os.File
	p *Person
	err error
	next bool
	scrolled bool
	t json.Token
}

func (r *PersonScanner) Count() int {
	return r.count
}

func (r *PersonScanner) Scan() (*Person, error) {
	return r.p, r.err
}

func (r *PersonScanner) Next() bool {
	defer func() {
		if r.next == false {
			r.file.Close()
		}
	}()

	if !r.scrolled {
		r.next = r.scroll()
		r.scrolled = r.next
	}

	if err := r.dec.Decode(&r.p); err != nil {
		r.next = false
		return r.next
	}
	r.count++

	r.next = true
	return r.next
}

func (r *PersonScanner) scroll() bool {
	for r.dec.More() {
		if r.t, r.err = r.dec.Token(); r.err != nil {
			return false
		}
		if r.t == "persons" {
			if r.t, r.err = r.dec.Token(); r.err != nil {
				return false
			}
			break;
		}
	}
	return true
}

type Person struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
}
