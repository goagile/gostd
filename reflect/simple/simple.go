package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID   int
	Name string
}

func main() {
	var i int = 2
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println(t)
	fmt.Println(v)

	var u = &User{ID: 234, Name: "AAA"}
	t = reflect.TypeOf(u)
	v = reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Printf("%+v\n", v)
	e := v.Elem()
	fmt.Println("e.NumField -> ", e.NumField())
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		fmt.Println("e.Kind() -> ", e.Kind())
		fmt.Println("f -> ", f)
	}
}
