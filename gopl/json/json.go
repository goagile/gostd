package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var result interface{}
	str := []byte(`{"x": 2}`)

	err := json.Unmarshal(str, &result)
	if err != nil {
		panic("AAA")
	}

	fmt.Println(result)
}
