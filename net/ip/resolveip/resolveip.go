package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "usage: ./resolveip hostname")
		fmt.Fprintln(os.Stderr, "example: ./resolveip google.ru")
		os.Exit(1)
	}
	host := args[0]
	ip, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot resolve ip for host", host)
		os.Exit(1)
	}
	fmt.Println("resolved ip: ", ip, " for host", host)
}
