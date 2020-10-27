package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const prefix = "http://"

func main() {
	for _, arg := range os.Args[1:] {
		url := arg
		if !strings.HasPrefix(prefix, url) {
			url = fmt.Sprintf("%s%s", prefix, url)
		}

		resp, err := http.Get(url)
		if err != nil {
			log.Println("Get", url, err)
			continue
		}

		log.Println("Status", resp.StatusCode, resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			log.Println("Copy", url, err)
			continue
		}

		resp.Body.Close()
	}
}
