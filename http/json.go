package main

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := loc{}

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error while read r.Body", err)
	}

	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error", err)
	}

	log.Printf("Recieved: %v\n", location)

	z := loc{
		Lat: 22,
		Lon: 44,
	}

	jsonz, err := json.Marshal(z)
	if err != nil {
		fmt.Fprintf(w, "Error in jsonz Marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonz)
}

func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8081", nil)
}

func client() {
	locJson, err := json.Marshal(loc{Lat: 2, Lon: 4})
	req, err := http.NewRequest("POST", "http://localhost:8081", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	go server()
	client()
}
