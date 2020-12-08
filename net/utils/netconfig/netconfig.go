package main

import (
	"log"
	"net"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile)
	for _, nf := range netInterfaces() {
		printInterface(nf)
	}
}

func netInterfaces() []net.Interface {
	infs, err := net.Interfaces()
	if err != nil {
		log.Fatal("net.Interfaces", err)
	}
	return infs
}

func printInterface(nf net.Interface) {
	println(nf.Name)
	printAddressesFor(nf)
	printHardwareAddressFor(nf)
	println("  ", nf.Flags.String())
	println("   MTU:", nf.MTU)
	println()
}

func printHardwareAddressFor(i net.Interface) {
	hw := i.HardwareAddr.String()
	if "" == strings.TrimSpace(hw) {
		return
	}
	println("   HWaddr:", hw)
}

func printAddressesFor(nf net.Interface) {
	i, err := net.InterfaceByName(nf.Name)
	if err != nil {
		log.Println("net.InterfaceByName", err)
		return
	}
	addrs, err := i.Addrs()
	if err != nil {
		log.Println("i.Addrs", err)
		return
	}
	for _, a := range addrs {
		printAddress(a)
	}
}

func printAddress(a net.Addr) {
	println("  ", a.Network(), a.String())
}
