package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// New client to TCP Clock with NetCat app:
// $ nc 127.0.0.1 8081

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TCP Clock: 127.0.0.1:8081")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// parallel
		go handle(conn)

	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		t := time.Now()
		s := fmt.Sprintf("%02v:%02v:%02v\n", t.Hour(), t.Minute(), t.Second())
		io.WriteString(conn, s)
		time.Sleep(1 * time.Second)
	}
}
