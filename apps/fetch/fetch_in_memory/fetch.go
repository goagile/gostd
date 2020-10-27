package main

import (
	"fmt"
	"io/ioutil"
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
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ReadAll", url, err)
			continue
		}
		resp.Body.Close()
		fmt.Println(b)
	}
}
