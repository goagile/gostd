package main

import (
	"fmt"
	// "io"
	// "io/ioutil"
	"os"
	"log"
	"encoding/xml"
)

func main() {

	s, err := NewPersonScanner("./example_4.xml")
	if err != nil {
		log.Fatalf("NewPersonScanner: %v\n", err)
	}	

	for s.Next() {
		p, err := s.Scan()
		if err != nil {
			log.Fatalf("Scan: %v\n", err)	
		}
		fmt.Println(p)
	}

}


func NewPersonScanner(filename string) (*PersonScanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Open: %v\n", err)
	}
	s := &PersonScanner{
		dec: xml.NewDecoder(f),
		f: f,
		next: true,
	}
	return s, nil
}

type PersonScanner struct {
	dec *xml.Decoder
	p *Person
	err error
	t xml.Token
	f *os.File
	next bool
}

func (s *PersonScanner) Scan() (*Person, error) {
	return s.p, s.err
}

func (s *PersonScanner) Next() bool {
	defer func() {
		if s.next == false {
			s.f.Close()
		}
	}()

	for {
		s.t, s.err = s.dec.Token()
		if s.t == nil || s.err != nil {
			s.next = false
			return s.next
		}
		switch e := s.t.(type) {
		default:
			continue
		case xml.StartElement:
			if e.Name.Local == "person" {
				s.err = s.dec.DecodeElement(&s.p, &e)
				if s.err != nil {
					s.next = false
					return s.next
				}
				s.next = true
				return s.next
			}
		}
	}

	s.next = true
	return s.next
}

type Person struct {
	Name string `xml:"name"`
	Age int `xml:"age"`
}

func (p *Person) String() string {
	return fmt.Sprintf("Person %q is %v years old", p.Name, p.Age)
}
