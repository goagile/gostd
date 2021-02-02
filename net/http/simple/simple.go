package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello: " + r.URL.Path + "\r\n"))
	})
	port := ":8080"
	fmt.Println("Server starts at ", port)
	http.ListenAndServe(port, nil)
}
