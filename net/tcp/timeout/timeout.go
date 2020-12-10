package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("usage: ./timeout port, examle: ./timeout 8081")
	}
	port := fmt.Sprintf(":%v", args[0])

	addr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		log.Fatalln("ResolveTCPAddr", err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln("ListenTCP", err)
	}

	for {
		c, err := l.AcceptTCP()
		if err != nil {
			log.Println("Accept", err)
			continue
		}

		t := time.Now().Add(1 * time.Second)
		if err := c.SetDeadline(t); err != nil {
			log.Println("SetDeadline", err)
			continue
		}

		now := time.Now().Format(time.RFC3339)
		fmt.Fprintln(c, now)
		c.Close()
	}
}
