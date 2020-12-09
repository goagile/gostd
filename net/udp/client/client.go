package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	addr := address()
	log.Println("UDP Client: ", addr)

	udp := resolve(addr)
	conn := client(udp)
	defer conn.Close()

	go readLoop(conn)
	writeLoop(conn)
}

func address() string {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("usage: ./server host port")
	}

	return net.JoinHostPort(args[0], args[1])
}

func resolve(addr string) *net.UDPAddr {
	udpa, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		log.Println("ResolveUDPAddr", err)
	}

	return udpa
}

func client(udpa *net.UDPAddr) *net.UDPConn {
	conn, err := net.DialUDP("udp4", nil, udpa)
	if err != nil {
		log.Println("DialUDP", err)
	}

	return conn
}

func writeLoop(conn *net.UDPConn) {
	for {
		if err := write(conn); err != nil {
			log.Println(err)
			continue
		}
	}
}

func readLoop(conn *net.UDPConn) {
	for {
		if err := read(conn); err != nil {
			log.Println(err)
			continue
		}
	}
}

func write(conn *net.UDPConn) error {
	fmt.Print(">> ")

	r := bufio.NewReader(os.Stdin)
	b, err := r.ReadBytes('\n')
	if err != nil {
		return fmt.Errorf("ReadString %v", err)
	}
	quit(b)

	b = append(b, "\n"...)
	n, err := conn.Write(b)
	if err != nil {
		return fmt.Errorf("Write %v", err)
	}
	log.Println("Write ", n, " bytes")

	return nil
}

func quit(b []byte) {
	switch strings.TrimSpace(string(b)) {
	case "q", "quit":
		fmt.Println("Bye!")
		os.Exit(0)
	}
}

func read(conn *net.UDPConn) error {
	b := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(b)
	if err != nil {
		return fmt.Errorf("ReadFromUDP %v %v", addr, err)
	}

	log.Println("Read ", n, " bytes from ", addr)
	log.Println(string(b))

	return nil
}
