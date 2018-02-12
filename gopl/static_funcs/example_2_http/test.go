package main


import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hi, I am a %s", r.URL.Path[1:])
}


func main() {
	fmt.Println("Server runs on port: %s", "8000")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
