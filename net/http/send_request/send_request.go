package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

var (
	data = `{"name":"Go"}`
	port = ":8080"
	url  = "http://127.0.0.1:8080/raw"
)

func main() {
	go send()
	http.HandleFunc("/raw", raw)
	fmt.Println("Listen at", port)
	http.ListenAndServe(port, nil)
}

func send() {
	time.Sleep(2 * time.Second)
	client := http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{}).Dial,
		},
	}
	body := bytes.NewBufferString(data)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		fmt.Println("NewRequest", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do", err)
		return
	}
	fmt.Printf("Send: %v\n", data)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return
	}
	fmt.Printf("Send Response: %v\n", string(b))
}

func raw(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ReadAll", err)
		return
	}
	fmt.Printf("Receive: %v\n", string(b))
	fmt.Fprintln(w, "OK")
}
