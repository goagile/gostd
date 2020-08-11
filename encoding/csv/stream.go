package main

import (
	"fmt"
	"encoding/csv"
	"strings"
	"strconv"
	"log"
	"io"
)

func main() {
	s := strings.NewReader(
		 "name,age\n"+
		 "AAAA,32\n" + 
		 "BBBB,34")

	r := NewPersonScanner(s)

	for r.Next() {
		p, err := r.Scan()
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		fmt.Println(p)
	}
}

func NewPersonScanner(r io.Reader) *PersonScanner {
	s := &PersonScanner{
		r: csv.NewReader(r),
		row: make([]string, 0),
		next: true,
	}
	s.skipHead()
	return s
}

type PersonScanner struct {
	r *csv.Reader
	err error
	row []string
	p *Person
	next bool
}

func (s *PersonScanner) skipHead() {
	s.r.Read()
}

func (s *PersonScanner) Scan() (*Person, error) {
	return s.p, s.err
}

func (s *PersonScanner) Next() bool {
	s.row, s.err = s.r.Read()
	if s.err != nil {
		s.next = false
		return s.next 
	}

	age, _ := strconv.Atoi(s.row[1])
	s.p = &Person{
		Name: s.row[0],
		Age: age,
	}

	s.next = true
	return s.next 
}

type Person struct {
	Name string
	Age int
}
