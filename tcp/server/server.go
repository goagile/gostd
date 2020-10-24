package main

import (
	"bufio"
	"flag"
	"log"
	"net"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "127.0.0.1:8000", "")

	flag.Parse()

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listen", addr)

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("Accept:", err)
			continue
		}

		log.Println("Accepted:", c.RemoteAddr())
		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	for {
		r := bufio.NewReader(c)
		s, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(s)
	}
}
