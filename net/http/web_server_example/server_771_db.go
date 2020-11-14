package main

import (
	"fmt"
	"net/http"
)

type price float32

func (p price) String() string {
	return fmt.Sprintf("$%.2f", p)
}

type database map[string]price

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%q: %v\n", item, price)
	}
}

func main() {
	host := "localhost"
	port := "8081"
	link := fmt.Sprintf("%s:%s", host, port)

	db := database{"socks": 20, "shoes": 3}

	fmt.Println("Server start at: ", link)
	http.ListenAndServe(link, db)
}
