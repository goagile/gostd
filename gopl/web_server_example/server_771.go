package main

import (
	"fmt"
	"net/http"
)

func main() {
	host := "localhost"
	port := "8081"
	link := fmt.Sprintf("%s:%s", host, port)

	fmt.Println("Server start at: ", link)

	http.HandleFunc("/", handler)
	http.ListenAndServe(link, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello</h1>")
}
