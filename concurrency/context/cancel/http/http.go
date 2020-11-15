package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer log.Println("Defer")
		log.Println("Handle Request")
		ctx := r.Context()
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				log.Fatal("Client: ", err)
			}
			log.Println("Done")
		case <-time.After(1 * time.Second):
			log.Fatal("Timeout")
		}
	})
	log.Println("Start ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
