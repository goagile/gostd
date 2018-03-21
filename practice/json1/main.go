package main

import (
	"encoding/json"
	"fmt"
)

type ErrLog struct {
	messages []string
}

func (e *ErrLog) Messages() []string {
	return e.messages
}

func (e *ErrLog) Add(message string) {
	e.messages = append(e.messages, message)
}

type PersonResult struct {
	Persons []Person `json:persons`
}

func (r *PersonResult) Validate(errLog *ErrLog) bool {
	result := true
	for _, p := range r.Persons {
		result = p.Validate(errLog)
	}
	return result
}

type Person struct {
	FirstName NotEmptyName `json:first_name`
	Education []Education  `json:education`
}

func (p *Person) Validate(errLog *ErrLog) bool {
	// return p.FirstName.Validate(errLog, "first_name")
	if !p.FirstName.Validate(errLog) {
		return false
	}
	return true
}

type Education struct {
	Start string `json:start`
	End   string `json:end`
	Name  string `json:name`
}

type NotEmptyName struct {
	Value string
}

func (n NotEmptyName) Validate(errLog *ErrLog) bool {
	if n.Value == "" {
		errLog.Add("пустое поле")
		return false
	}
	return true
}

func main() {
	f1()
	// f2()
	// f3()
	// testNotEmptyName()
}

func testNotEmptyName() {
	firstName := NotEmptyName{""}

	errLog := new(ErrLog)
	isValid := firstName.Validate(errLog)
	fmt.Printf("isValid: %v\n", isValid)

	fmt.Println("Сообщения:")
	for _, msg := range errLog.Messages() {
		fmt.Println(msg)
	}
}

func f3() {
	person := Person{}

	errLog := new(ErrLog)
	isValid := person.Validate(errLog)
	fmt.Printf("isValid: %v\n", isValid)

	fmt.Println("Сообщения:")
	for _, msg := range errLog.Messages() {
		fmt.Println(msg)
	}
}

func f1() {
	fmt.Println("f1")
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

	var personResult PersonResult

	if err := json.Unmarshal(byt, &personResult); err != nil {
		panic(err)
	}
	fmt.Println("Результат маршалинга:")
	fmt.Printf("%+v\n", personResult)

	errLog := new(ErrLog)
	isValid := personResult.Validate(errLog)
	fmt.Printf("isValid: %v\n", isValid)

	fmt.Println("Сообщения:")
	for _, msg := range errLog.Messages() {
		fmt.Println(msg)
	}
}

func f2() {
	byt := []byte(`{
		"persons": [
			{
				"first_name": null,
				"education": [
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
