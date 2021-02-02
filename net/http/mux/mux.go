package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	server := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	fmt.Println("Server at ", port)
	server.ListenAndServe()
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home", r.URL.String())
}
