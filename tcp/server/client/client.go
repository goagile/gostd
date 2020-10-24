package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var addr string
	flag.StringVar(&addr, "a", "addr", "127.0.0.1:8000")

	flag.Parse()

	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("net.Dial", err)
	}

	for {
		fmt.Print("> ")
		r := bufio.NewReader(os.Stdin)
		cmd, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		n, err := io.WriteString(c, cmd+"\n")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%v send %v bytes\n", cmd, n)
	}
}
