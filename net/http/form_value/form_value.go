package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)
	srv := &http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Println("Serve at ", port)
	srv.ListenAndServe()
}

func handle(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	p := fmt.Sprintf("You are send param: %v\n", param)
	fmt.Println(p)
	fmt.Fprintf(w, p)

	key := r.FormValue("key")
	k := fmt.Sprintf("You are send key: %v\n", key)
	fmt.Println(k)
	fmt.Fprintf(w, k)
}
