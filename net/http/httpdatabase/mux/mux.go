package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	host = "127.0.0.1"
	port = 8080
	src  = fmt.Sprintf("%v:%v", host, port)
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", home)
	m.HandleFunc("/list", list)
	log.Fatal(http.ListenAndServe(src, m))
}

func home(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OKI, OKI")
}

func list(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "AAA")
}
