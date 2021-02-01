package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	a := conn.RemoteAddr()
	fmt.Println(a, "Connencted")
	fmt.Fprintf(conn, "You are %v", a)
	s := bufio.NewScanner(conn)
	for s.Scan() {
		t := strings.TrimSpace(s.Text())
		if t == "" {
			continue
		}
		if t == "Quit" || t == "Exit" {
			fmt.Fprintf(conn, "Bye, Bye!\r\n")
			fmt.Println(a, "Exited")
			return
		}
		fmt.Println(a, t)
	}
	fmt.Fprintf(conn, "\r\n")
	fmt.Println(a, "Disconnected")
}
