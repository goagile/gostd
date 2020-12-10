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
		log.Println("usage ./lookupport ip http")
	}

	network, service := args[0], args[1]

	port, err := net.LookupPort(network, service)
	if err != nil {
		log.Println("LookupPort", err)
	}

	fmt.Println("port: ", port, " for ", network, service)
}
