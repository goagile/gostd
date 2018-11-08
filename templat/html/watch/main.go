package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

type Post struct {
	Title string
	Body  string
}

var (
	watch bool
	post  *Post
	t     *template.Template
)

func init() {
	post = &Post{
		Title: "Hello  Templating",
		Body:  "Templating is amazing.<script>alert('NO XSS')</script>",
	}
	t = parseTemplate("index.html")
}

func parseTemplate(filenames ...string) *template.Template {
	return template.Must(template.ParseFiles(filenames...))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if watch {
		t = parseTemplate("index.html")
	}
	t.Execute(w, post)
}

func main() {
	// читаем параметры запуска
	flag.BoolVar(&watch, "w", false, "-w watch: следит за изменением шаблонов")
	flag.Parse()

	// запускаем сервер
	http.HandleFunc("/", handler)
	log.Println("Server was run at port 3000")
	http.ListenAndServe(":3000", nil)
}
