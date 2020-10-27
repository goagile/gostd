package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Get", url, err)
			continue
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Println("Copy", url, err)
			continue
		}
		err = resp.Body.Close()
		if err != nil {
			log.Println("Close", url, err)
			continue
		}
	}
}
