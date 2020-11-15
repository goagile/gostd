package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	url     = "http://google.com"
	timeout = 180 * time.Millisecond
)

func main() {
	fmt.Println("timeout", timeout)

	ctx := context.Background()
	ctx, fin := context.WithTimeout(ctx, timeout)
	defer fin()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure", err)
		return
	}
	defer res.Body.Close()
	fmt.Println("Success", res.StatusCode)
}
