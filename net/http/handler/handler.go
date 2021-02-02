package main

import (
	"fmt"
	"net/http"
)

// Handler -
type Handler struct {
	Name string
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handelr:%v, %v\n", h.Name, r.URL.String())
}

func main() {
	http.Handle("/", &Handler{Name: "Home"})
	http.Handle("/page", &Handler{Name: "Page"})
	http.Handle("/pages/", &Handler{Name: "Pages"})
	port := ":8080"
	fmt.Println("Listen at ", port)
	http.ListenAndServe(port, nil)
}
