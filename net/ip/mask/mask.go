package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("need ip address")
		os.Exit(1)
	}

	s := args[0]
	ip := net.ParseIP(s)
	if ip == nil {
		fmt.Println("cannot parse ip address")
		os.Exit(1)
	}
	fmt.Printf("IP(%T): %v\n", ip, ip)

	dmask := ip.DefaultMask()
	fmt.Printf("\nDefaultMask(%T): %v \n", dmask, dmask)

	ones, bits := dmask.Size()
	fmt.Printf("ones: %v bits: %v\n", ones, bits)

	network := ip.Mask(dmask)
	fmt.Printf("Network(%T): %v\n", network, network)

	mask := net.IPv4Mask(123, 11, 0, 0)
	fmt.Printf("\nMask(%T): %v \n", mask, mask)

	ones, bits = mask.Size()
	fmt.Printf("ones: %v bits: %v\n", ones, bits)

	network = ip.Mask(mask)
	fmt.Printf("Network(%T): %v\n", network, network)
}
