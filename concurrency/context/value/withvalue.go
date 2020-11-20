package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	wrapped := wrap(home)
	http.HandleFunc("/", wrapped)
	log.Println("Start ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home")
	ctx := r.Context()
	id := ctx.Value("id")
	if id == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "not found id")
		return
	}
	log.Println("id", id)
	r.URL.Query().Set("id", id.(string))
	w.WriteHeader(http.StatusOK)
	return
}

func wrap(f http.HandlerFunc) http.HandlerFunc {
	log.Println("Wrap Home")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Set id")
		ctx := context.WithValue(r.Context(), "id", "HELLO")
		r = r.WithContext(ctx)
		f(w, r)
	}
}
