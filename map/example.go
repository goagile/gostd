package main

import (
	"fmt"
)

type I interface{}
type D map[string]map[string]map[string]I

func main() {

	d := D{
		"a": {
			"b": {
				"d1": 224323,
				"d2": 2424,
			},
		},
	}

	fmt.Println(d)
}
