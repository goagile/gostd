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
		log.Fatalln("usage ./lookupcname host")
	}

	host := args[0]

	cname, err := net.LookupCNAME(host)
	if err != nil {
		log.Fatalln("LookupCNAME", err)
	}

	fmt.Println("cname: ", cname, " for host ", host)
}
