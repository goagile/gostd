package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	done := make(chan struct{}, 2)

	m1src := "localhost:8081"
	m1 := http.NewServeMux()
	m1.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic m1", err)
				done <- struct{}{}
			}
		}()
		select {
		case <-signals:
			done <- struct{}{}
		default:
			fmt.Fprint(w, "m1")
		}
	})
	go func() {
		select {
		case <-signals:
			done <- struct{}{}
		default:
			fmt.Println("listen m1")
			if err := http.ListenAndServe(m1src, m1); err != nil {
				log.Fatal(err)
				done <- struct{}{}
			}
		}
	}()

	m2src := "localhost:8082"
	m2 := http.NewServeMux()
	m2.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("panic m2", err)
				done <- struct{}{}
			}
		}()
		select {
		case <-signals:
			done <- struct{}{}
		default:
			panic("sss")
		}
	})
	go func() {
		select {
		case <-signals:
			done <- struct{}{}
		default:
			fmt.Println("listen m2")
			if err := http.ListenAndServe(m2src, m2); err != nil {
				log.Fatal(err)
				done <- struct{}{}
			}
		}
	}()

	<-done
}
