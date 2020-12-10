package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("usage: ./lookuphost host, example: ./lookuphost google.ru")
	}

	host := args[0]
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatalln("LookupHost", err)
	}

	fmt.Println("lookup ", host)
	for _, addr := range addrs {
		fmt.Println("\t", addr)
	}
}
