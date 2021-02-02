package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Home")
		})
	http.HandleFunc("/page",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.URL.String())
		})
	http.HandleFunc("/pages/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, r.URL.String())
		})
	port := ":8080"
	fmt.Println("Server start at ", port)
	http.ListenAndServe(port, nil)
}
