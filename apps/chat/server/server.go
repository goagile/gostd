package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	addr = "127.0.0.1:8081"
	msgs = make(chan Message)
)

type Message struct {
	Addr string
	Body []byte
}

func main() {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go listenConnections(l)
	go listenStopServer()

	listenMessages()
}

func listenConnections(l net.Listener) {
	log.Println("Start on", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()
	log.Printf("(!) %v connected\n", name)
	if _, err := fmt.Fprintf(conn, "Hello, you are %v\n", name); err != nil {
		log.Println(err)
		return
	}

	defer func(name string) {
		conn.Close()
		log.Printf("(!) %v leaved\n", name)
	}(name)

	for {
		b, err := bufio.NewReader(conn).ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		msgs <- Message{name, b}
	}
}

func listenStopServer() {
	for {
		b, err := bufio.NewReader(os.Stdin).ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Println(err)
			return
		}
		if "q\n" == string(b) {
			return
		}
	}
}

func listenMessages() {
	for msg := range msgs {
		bs := bytes.Split(msg.Body, []byte(" "))
		if len(bs) < 2 {
			log.Printf("%s: %s", msg.Addr, msg.Body)
			continue
		}
		cmd := string(bs[0])
		tail := bs[1]
		if "REG" == cmd {
			log.Println("RRRRRRRRRRRRR")
			log.Println(string(tail))
		}
		log.Printf("%s: %s", msg.Addr, msg.Body)
	}
}
