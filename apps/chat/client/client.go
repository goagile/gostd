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

var (
	name string
	host string
	port int
)

func main() {
	flag.StringVar(&host, "h", "127.0.0.1", "host of chat server")
	flag.IntVar(&port, "p", 8001, "port of chat server")
	flag.StringVar(&name, "n", "", "name of client")
	flag.Parse()
	addr := fmt.Sprintf("%v:%v", host, port)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go sendClientName(conn, name)
	go listenServer(conn)

	listenStdin(conn)
}

func sendClientName(conn net.Conn, name string) {
	var msg []byte
	msg = append(msg, fmt.Sprintf("REG @%v\n", name)...)
	if _, err := conn.Write(msg); err != nil {
		log.Println(err)
		return
	}
}

func listenServer(conn net.Conn) {
	defer func() {
		conn.Close()
		log.Println("Good by!")
	}()
	for {
		msg, err := bufio.NewReader(conn).ReadBytes('\n')
		if err == io.EOF {
			log.Println("Server disconnected")
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%s", msg)
	}
}

func listenStdin(conn net.Conn) {
	for {
		msg, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println("ReadBytes", err)
			return
		}
		if "q\n" == string(msg) {
			log.Println("Q")
			return
		}
		if _, err := conn.Write(msg); err != nil {
			log.Println("conn.Write", err)
			return
		}
	}
}
