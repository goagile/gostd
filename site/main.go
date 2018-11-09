package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"trygo/site/core"
	"trygo/site/core/search"

	"github.com/bmizerany/pat"
)

var (
	PostTemplate = template.Must(template.ParseFiles(
		path.Join("templates", "layout.html"),
		path.Join("templates", "post.html"),
	))

	ErrorTemplate = template.Must(template.ParseFiles(
		path.Join("templates", "layout.html"),
		path.Join("templates", "error.html"),
	))

	BlevePath = "example.bleve"
)

func main() {
	fs := http.FileServer(http.Dir("./public/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	mux := pat.New()
	mux.Get("/:page", http.HandlerFunc(handler))
	mux.Get("/:page/", http.HandlerFunc(handler))
	mux.Get("/", http.HandlerFunc(handler))
	http.Handle("/", mux)

	http.HandleFunc("/search/", searchhandler)

	fmt.Println("Server was running at :3000")
	http.ListenAndServe(":3000", nil)
}

func searchhandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("interna server err:", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("read request body err:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		log.Println("unmarshal request body err:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	searchstring, ok := data["searchstring"].(string)
	if !ok {
		log.Println("searchstring convert err")
		http.Error(w, http.StatusText(400), 400)
		return
	}

	postShortInfo, err := search.FindPostShortInfo(BlevePath, searchstring)
	if err != nil {
		log.Println("find post short info err:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	byt, err := json.Marshal(postShortInfo)
	if err != nil {
		log.Println("post short info marshal err:", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.WriteHeader(200)
	w.Write(byt)
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	page := params.Get(":page")
	p := path.Join("posts", page)
	var post_md string
	if page != "" {
		post_md = p + ".md"
	} else {
		post_md = p + "/index.md"
	}
	post, status, err := core.LoadPost(post_md)
	if err != nil {
		log.Println(err)
		errorHandler(w, r, status)
		return
	}
	if err := PostTemplate.ExecuteTemplate(w, "layout", post); err != nil {
		log.Println(err)
		errorHandler(w, r, status)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if err := ErrorTemplate.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Error":  http.StatusText(status),
		"Status": status,
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}
}
