package main

import (
	"fmt"
	"net"
)

const src = "scanme.nmap.org"

func main() {
	s := NewPortScanner(100)
	go s.Scan()
	for p := range s.opened {
		fmt.Println(p)
	}
}

func NewPortScanner(n int) *PortsScanner {
	s := &PortsScanner{
		ports:  make(chan int, n),
		opened: make(chan int),
	}
	return s
}

type PortsScanner struct {
	ports  chan int
	opened chan int
}

func (s *PortsScanner) Scan() chan int {
	go s.fillports()
	for p := range s.ports {
		go s.scanport(p)
	}
	return s.opened
}

func (s *PortsScanner) fillports() {
	for p := 1; p < cap(s.ports); p++ {
		s.ports <- p
	}
}

func (s *PortsScanner) scanport(p int) {
	c, err := net.Dial("tcp", addr(p))
	if err != nil {
		return
	}
	c.Close()
	s.opened <- p
}

func addr(p int) string {
	return fmt.Sprintf("%v:%v", src, p)
}
