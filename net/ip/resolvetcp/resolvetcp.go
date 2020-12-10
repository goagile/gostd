package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatalln("usage ./resolvetcp tcp www.google.ru:80")
	}

	network, addr := args[0], args[1]
	tcpaddr, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		log.Fatalln("ResolveTCPAddr", err)
	}
	fmt.Println("network: ", network, "addr: ", addr)
	fmt.Println("tcpaddr.IP: ", tcpaddr.IP)
	fmt.Println("tcpaddr.Port: ", tcpaddr.Port)
	fmt.Println("tcpaddr.Zone: ", tcpaddr.Zone)
	fmt.Println("tcpaddr.Network(): ", tcpaddr.Network())
}
