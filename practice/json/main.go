package main

import (
	"encoding/json"
	"fmt"
)

// PersonResult dsfsdf sd fsdf
type PersonResult struct {
	Persons []Person `json:persons`
}

type Person struct {
	FirstName string      `json:first_name`
	Education []Education `json:education`
}

type Education struct {
	Start string `json:start`
	End   string `json:end`
	Name  string `json:name`
}

func main() {
	byt := []byte(`{
		"persons": [
			{
				"first_name": "Петр",
				"education": [
					{
						"start": "2001",
						"end": "2007",
						"name": "МГУ"
					}
				]
			}
		]
	}`)

	var dat PersonResult

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", dat)
}
