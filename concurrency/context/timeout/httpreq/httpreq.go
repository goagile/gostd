package main

import (
	"context"
	"fmt"
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
	ctx, _ = context.WithTimeout(ctx, timeout)

	req, _ := http.NewRequest("GET", url, nil)
	req = req.WithContext(ctx)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Failure", err)
		return
	}
	fmt.Println("Success", res.StatusCode)
}
