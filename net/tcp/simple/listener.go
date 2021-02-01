package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := ":8080"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Listen", err)
	}
	fmt.Println("Listen on", port)
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Accept", err)
		}
		go handle(c)
	}
}

func handle(conn net.Conn) {
	fmt.Println("Connencted: ", conn.RemoteAddr())
	fmt.Fprintf(conn, "OK")
}
