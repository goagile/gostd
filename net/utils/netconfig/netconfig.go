package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)
	for _, i := range netInterfaces() {
		printInterface(i)
	}
}

func netInterfaces() []net.Interface {
	is, err := net.Interfaces()
	if err != nil {
		log.Fatal("net.Interfaces", err)
	}
	return is
}

func printInterface(i net.Interface) {
	fmt.Println(i.Name)
	printAddressesFor(i)
	printHardwareAddressFor(i)
	fmt.Println("  ", i.Flags.String())
	fmt.Println("   MTU:", i.MTU)
	fmt.Println()
}

func printHardwareAddressFor(i net.Interface) {
	hw := i.HardwareAddr.String()
	if "" == strings.TrimSpace(hw) {
		return
	}
	fmt.Println("   HWaddr:", hw)
}

func printAddressesFor(i net.Interface) {
	n, err := net.InterfaceByName(i.Name)
	if err != nil {
		log.Println("net.InterfaceByName", err)
		return
	}
	addrs, err := n.Addrs()
	if err != nil {
		log.Println("n.Addrs", err)
		return
	}
	for _, a := range addrs {
		printAddress(a)
	}
}

func printAddress(a net.Addr) {
	fmt.Println("  ", a.Network(), a.String())
}
