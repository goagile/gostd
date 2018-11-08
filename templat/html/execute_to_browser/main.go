package main

import (
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Title string
	Body  string
}

var post = &Post{
	Title: "Hello  Templating",
	Body:  "Templating is amazing.<script>alert('NO XSS')</script>",
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, post)
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server was run at port 3000")
	http.ListenAndServe(":3000", nil)
}
