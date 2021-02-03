package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var html = []byte(`
	<html>
	<body>
		<form action="/upload" method="post" enctype="multipart/form-data">
			Image: <input type="file" name="image"/>
			<input type="submit" value="Up"/>
		</form>
	</body>
	</html>
`)

func main() {
	port := ":8080"
	http.HandleFunc("/", home)
	http.HandleFunc("/upload", upload)
	fmt.Println("Listen at", port)
	http.ListenAndServe(port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write(html)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1024)
	u, h, err := r.FormFile("image")
	if err != nil {
		fmt.Println("FormFile", err)
		fmt.Fprintln(w, "FormFile", err)
		return
	}
	fmt.Println(h.Filename)
	b, err := ioutil.ReadAll(u)
	if err != nil {
		fmt.Println("ReadAll", err)
		fmt.Fprintln(w, "ReadAll", err)
		return
	}
	ioutil.WriteFile(h.Filename, b, 0644)
	fmt.Fprintf(w, "Uploaded %v OK", h.Filename)
}
