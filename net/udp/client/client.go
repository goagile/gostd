package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("usage: ./client host port")
	}
	a := net.JoinHostPort(args[0], args[1])

	raddr, err := net.ResolveUDPAddr("udp4", a)
	if err != nil {
		log.Fatalln("ResolveUDPAddr", err)
	}

	c, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		log.Fatalln("DialUDP", err)
	}
	defer c.Close()

	go readLoop(c)

	writeLoop(c)
}

func readLoop(c *net.UDPConn) {
	for {
		b := make([]byte, 1024)
		n, addr, err := c.ReadFromUDP(b)
		if err != nil {
			log.Println("ReadFromUDP", err)
			continue
		}
		fmt.Printf("\n<-- read %v bytes from server %v: %v\n", n, addr, string(b))
	}
}

func writeLoop(c *net.UDPConn) {
	for {
		fmt.Print(">>> ")
		b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err != nil {
			log.Println("ReadBytes", err)
			continue
		}
		b = bytes.Trim(b, "\n")
		if len(b) == 0 {
			continue
		}
		n, err := c.Write(b)
		if err != nil {
			log.Println("Write", err)
			continue
		}
		a := c.RemoteAddr().String()
		fmt.Printf("--> send %v bytes to server %v: %v\n", n, a, string(b))
	}
}
