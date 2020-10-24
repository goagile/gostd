package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const timeout = 10 * time.Second
const maxReadBuf = 4 // 4 byte

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
		c.SetDeadline(time.Now().Add(timeout))
		log.Println("Accepted:", c.RemoteAddr())
		go handle(c)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	for {
		var cmd string
		r := io.LimitReader(c, maxReadBuf)
		if _, err := fmt.Fscan(r, &cmd); err != nil {
			log.Println(err)
			break
		}
		log.Println(cmd)
		c.SetDeadline(time.Now().Add(timeout))
	}
}
