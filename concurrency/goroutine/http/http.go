package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	// run client
	go func() {
		time.Sleep(3 * time.Second)
		http.Get("http://127.0.0.1:8081")
	}()

	// run server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			log.Println("Hello")
		}()
		return
	})
	log.Println("Starting at :8081")
	http.ListenAndServe(":8081", nil)

}
