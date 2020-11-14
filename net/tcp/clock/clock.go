package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Clock was run: 127.0.0.1:8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	for {
		t := time.Now()
		s := fmt.Sprintf("%02v:%02v:%02v\n", t.Hour(), t.Minute(), t.Second())
		_, err := io.WriteString(conn, s)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
