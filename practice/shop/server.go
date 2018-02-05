package main

import (
	"fmt"
	"log"
	"net/http"
)

type database map[string]int

func (db database) List(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%q", r.URL.Query())
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

func main() {
	host := "localhost"
	port := "8081"
	link := fmt.Sprintf("%s:%s", host, port)
	
	db := database{"shoes": 20, "socks": 4}

	http.HandleFunc("/", db.List)

	fmt.Println("Server run at:", link)
	log.Fatal(http.ListenAndServe(link, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// for k, v := range r {
	// 	fmt.Fprintf(w, "%q: %q\n", k, v)		
	// }
	// fmt.Fprintf(w, "%s\n", "Hello")
}
