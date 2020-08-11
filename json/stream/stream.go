package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	"log"
	"encoding/json"
)

func main() {

	r, err := NewPersonReader("./test/example.json")
	if err != nil {
		log.Fatalf("NewPersonReader %v\n", err)
	}

	for r.NextPerson() {
		p, err := r.ScanPerson()
		if err != nil {
			log.Fatalf("ScanPerson err: %v\n", err)
		}
		fmt.Println(p)
	}

	fmt.Printf("Read %v persons.\n", r.Count())

}

func NewPersonReader(filename string) (*PersonReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := &PersonReader{
		dec: json.NewDecoder(f),
		next: true,
		scrolled: false,
	}
	return r, nil
}

type PersonReader struct {
	dec *json.Decoder
	count int
	file *os.File
	p *Person
	err error
	next bool
	scrolled bool
}

func (r *PersonReader) Count() int {
	return r.count
}

func (r *PersonReader) ScanPerson() (*Person, error) {
	return r.p, r.err
}

func (r *PersonReader) NextPerson() bool {
	defer func() {
		if r.next == false {
			r.file.Close()
		}
	}()

	if !r.scrolled {
		for r.dec.More() {
			t, err := r.dec.Token()
			if err != nil {
				r.next = false
				return r.next
			}
			if t == "persons" {
				fmt.Printf("%T %v\n", t, t)
				if _, err := r.dec.Token(); err != nil {
					r.next = false
					return r.next
				}
				break;
			}
		}
		r.scrolled = true
	}

	if err := r.dec.Decode(&r.p); err != nil {
		r.next = false
		return r.next
	}
	r.count++

	r.next = true
	return r.next
}

type Person struct {
	Name string `json:"name"`
	Age float64 `json:"age"`
}
