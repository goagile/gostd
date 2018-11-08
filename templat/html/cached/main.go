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

var (
	post *Post
	t    *template.Template
)

func init() {
	post = &Post{
		Title: "Hello  Templating",
		Body:  "Cached template",
	}
	t = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, post)
	})
	log.Println("Server was run at port 3000")
	http.ListenAndServe(":3000", nil)
}
