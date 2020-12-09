package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("usage: ./server host port")
	}
	a := net.JoinHostPort(args[0], args[1])
	fmt.Println("udp server listen: ", a)

	laddr, err := net.ResolveUDPAddr("udp4", a)
	if err != nil {
		log.Fatalln("ResolveUDPAddr", err)
	}

	c, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		log.Fatalln("ListenUDP", err)
	}
	defer c.Close()

	for {
		b := make([]byte, 1024)
		n, addr, err := c.ReadFromUDP(b)
		if err != nil {
			log.Println("ReadFromUDP", err)
			continue
		}
		fmt.Printf("\n<-- read %v bytes from client %v: %v\n", n, addr, string(b))

		fmt.Println("!!! start ... ")
		time.Sleep(3 * time.Second)
		fmt.Println("!!! end ... ")

		m, err := c.WriteToUDP(b, addr)
		if err != nil {
			log.Println("WriteToUDP", err)
			continue
		}
		fmt.Printf("--> send %v bytes to client %v: %v", m, addr, string(b))
	}
}
