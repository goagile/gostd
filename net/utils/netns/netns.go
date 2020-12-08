package main

import (
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	for _, domain := range domains(os.Args[1:]) {
		for _, host := range nameServers(domain) {
			log.Println(host)
		}
	}
}

func domains(args []string) []string {
	if len(args) == 0 {
		log.Fatalln("need domain name(s)")
	}
	return args
}

func nameServers(domain string) (hosts []string) {
	nameServers, err := net.LookupNS(domain)
	if err != nil {
		log.Fatalln("net.LookupNS", domain, err)
	}
	for _, ns := range nameServers {
		hosts = append(hosts, ns.Host)
	}
	return hosts
}
