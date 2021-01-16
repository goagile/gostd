package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int    `json:"user_id"`
	Name string `json:"user_name"`
}

func main() {
	u := &User{
		ID:   234,
		Name: "UserA",
	}

	v := reflect.ValueOf(u).Elem()
	fmt.Println("NumField: ", v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		t := v.Type().Field(i)
		fmt.Println("t.Name:", t.Name)
		fmt.Println("v.Field():", f)
		fmt.Println("t.Type().Kind():", t.Type.Kind())
		fmt.Println("t.Tag:", t.Tag)
	}
}
