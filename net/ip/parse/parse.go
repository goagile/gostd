package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("want ip address")
		os.Exit(1)
	}

	s := args[0]
	ip := net.ParseIP(s)
	if ip == nil {
		fmt.Println("cannot parse ip address", s)
		os.Exit(2)
	}

	fmt.Printf("parsed ip: %T %v\n", ip, ip)
	fmt.Println("\tTo4 ", ip.To4())
	fmt.Println("\tTo16 ", ip.To16())
	fmt.Println("\tIsLoopback ", ip.IsLoopback())
	fmt.Println("\tDefaultMask ", ip.DefaultMask())
}
