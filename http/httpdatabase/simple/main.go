package main

import (
	"fmt"
	"log"
	"net/http"
)

var srv = "127.0.0.1:8080"

func main() {
	db := DB{
		"shoes":  1200,
		"tshort": 800,
	}
	log.Printf("Listen on http://%v\n", srv)
	log.Fatal(http.ListenAndServe(srv, db))
}

type DB map[string]ruble

func (db DB) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for k, v := range db {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}
	case "/price":
		key := req.URL.Query().Get("key")
		price, ok := db[key]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "не найден товар %v\n", key)
		}
		fmt.Fprintf(w, "%v\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "не найден метод %v\n", req.URL.Path)
	}
}

type ruble int

func (r ruble) String() string {
	return fmt.Sprintf("%vР", int(r))
}
