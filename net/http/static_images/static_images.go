package main

import (
	"fmt"
	"net/http"
)

var html = `
	<html>	
		<body>
			<img src="/static/img.png"/>
		</body>
	</html>
	`

func main() {
	port := ":8080"
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", home)
	fmt.Println("Listen at", port)
	http.ListenAndServe(port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(html))
}
