package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/raw", raw)
	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}
	fmt.Println("Listen at ", port)
	srv.ListenAndServe()
}

type Obj struct {
	Name string `json:"name"`
}

func raw(w http.ResponseWriter, r *http.Request) {
	ww := io.MultiWriter(w, os.Stdout)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(ww, fmt.Sprintf("ReadAll %v", err))
		return
	}
	var o Obj
	if err := json.Unmarshal(b, &o); err != nil {
		fmt.Fprintln(ww, fmt.Sprintf("ReadAll %v", err))
		return
	}
	fmt.Fprintf(ww, "OK: %+v\n", o)
}
