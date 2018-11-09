package core

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/russross/blackfriday"
)

type Post struct {
	Title    string
	Body     template.HTML
	Markdown string
}

func LoadPost(md string) (Post, int, error) {
	info, err := os.Stat(md)
	if err != nil {
		if os.IsNotExist(err) {
			// файл не существует
			return Post{}, http.StatusNotFound, err
		}
	}
	if info.IsDir() {
		// не файл, а папка
		return Post{}, http.StatusNotFound, fmt.Errorf("dir")
	}
	fileread, err := ioutil.ReadFile(md)
	if err != nil {
		log.Println("read .md error:", err)
		return Post{}, http.StatusNotFound, nil
	}
	lines := strings.Split(string(fileread), "\n")
	title := string(lines[0])
	body := strings.Join(lines[1:len(lines)], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post := Post{title, template.HTML(body), body}
	return post, 200, nil
}
