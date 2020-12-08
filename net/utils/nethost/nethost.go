package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	for _, name := range os.Args[1:] {
		printAddrs(name, lookup(name))
	}
}

func lookup(addr string) []string {
	if ip := net.ParseIP(addr); ip != nil {
		return hosts(ip.String())
	}
	return ips(addr)
}

func hosts(ip string) []string {
	names, err := net.LookupAddr(ip)
	if err != nil {
		log.Fatalln("net.LookupAddr", err)
	}
	return names
}

func ips(host string) []string {
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatalln("net.LookupHost", err)
	}
	return addrs
}

func printAddrs(name string, addrs []string) {
	for _, addr := range addrs {
		fmt.Println(name, " -> ", addr)
	}
}
