package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("usage: ./head host:port, example: ./head google.ru:80")
	}

	s := args[0]
	addr, err := net.ResolveTCPAddr("tcp4", s)
	if err != nil {
		log.Fatalln("ResolveTCPAddr", s, err)
	}
	fmt.Println("Resolved ", addr)

	c, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln("DialTCP", err)
	}
	defer c.Close()

	m := []byte(`HEAD / HTTP/1.0\r\n\r\n`)
	fmt.Println("Write to ", addr, "\n", string(m))
	_, err = c.Write(m)
	if err != nil {
		log.Fatalln("Write", err)
	}

	b, err := ioutil.ReadAll(c)
	if err != nil {
		log.Fatalln("ReadAll", err)
	}
	fmt.Println(string(b))

	os.Exit(0)
}
