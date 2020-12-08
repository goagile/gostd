package main

import (
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
	println(i.Name)
	printAddressesFor(i)
	printHardwareAddressFor(i)
	println("  ", i.Flags.String())
	println("   MTU:", i.MTU)
	println()
}

func printHardwareAddressFor(i net.Interface) {
	hw := i.HardwareAddr.String()
	if "" == strings.TrimSpace(hw) {
		return
	}
	println("   HWaddr:", hw)
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
	println("  ", a.Network(), a.String())
}
