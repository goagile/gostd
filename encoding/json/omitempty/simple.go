package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID      int    `json:"user_id"`
	Name    string `json:"user_name"`
	Age     int    `json:"age,omitempty"`
	Comment string `json:"-"`
}

func main() {
	//
	// Marshal
	//
	a := &User{
		ID:      123,
		Name:    "A",
		Comment: "First",
	}
	ba, _ := json.Marshal(a)
	fmt.Println(string(ba))

	//
	// Unmarshal
	//
	bb := []byte(`{"user_id":543, "user_name":"B"}`)
	var b User
	json.Unmarshal(bb, &b)
	fmt.Printf("%+v", b)
}
