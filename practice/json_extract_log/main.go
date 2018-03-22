package main

import (
	"encoding/json"
	"fmt"
)

type ErrLog struct {
	messages []string
	prefix   string
}

func (e *ErrLog) Messages() []string {
	return e.messages
}

func (e *ErrLog) SetPrefix(prefix string) {
	e.prefix = prefix
}

func (e *ErrLog) Add(message string) {
	msg := fmt.Sprintf("%s: %s", e.prefix, message)
	e.messages = append(e.messages, msg)
}

type PersonResult struct {
	Persons []Person `json:persons`
}

func (r *PersonResult) Validate(errLog *ErrLog) bool {
	result := true
	for i, p := range r.Persons {
		errLog.SetPrefix(fmt.Sprintf("персона №%v", i))
		if !p.Validate(errLog) {
			result = false
		}
	}
	return result
}

type Person struct {
	FirstName NotEmptyName `json:"first_name"`
	Education []Education  `json:"education"`
}

func (p *Person) Validate(errLog *ErrLog) bool {
	result := true
	errLog.SetPrefix(`поле "first_name"`)
	if !p.FirstName.Validate(errLog) {
		result = false
	}
	return result
}

type Education struct {
	Start string `json:start`
	End   string `json:end`
	Name  string `json:name`
}

type NotEmptyName struct {
	Value string
}

func (n *NotEmptyName) String() string {
	return n.Value
}

func (n *NotEmptyName) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &n.Value)
}

func (n *NotEmptyName) Validate(errLog *ErrLog) bool {
	if n.Value == "" {
		errLog.Add("пустое поле")
		return false
	}
	return true
}

func main() {
	f1()
	// testNotEmptyName()
	// testNotEmptyNameUNMARSHAL()
	// testPersonUNMARSHAL()
}

func testPersonUNMARSHAL() {
	fmt.Println("testPersonUNMARSHAL")
	byt := []byte(`{"first_name": "Петрович"}`)

	var person Person

	if err := json.Unmarshal(byt, &person); err != nil {
		panic(err)
	}
	fmt.Println("Результат маршалинга:")
	fmt.Printf("\t%+v\n", person)

	// errLog := new(ErrLog)
	// person.Validate(errLog)
	//
	// fmt.Println("Сообщения:")
	// for _, msg := range errLog.Messages() {
	// 	fmt.Printf("\t%v\n", msg)
	// }
}

func testNotEmptyNameUNMARSHAL() {
	byt := []byte(`"HELLO"`)
	var dat NotEmptyName

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", dat)
}

func f1() {
	fmt.Println("f1")
	byt := []byte(`{
		"persons": [
			{
				"first_name": "FFDSD"
			}
		]
	}`)

	var personResult PersonResult

	if err := json.Unmarshal(byt, &personResult); err != nil {
		panic(err)
	}
	fmt.Println("Результат маршалинга:")
	fmt.Printf("\t%+v\n", personResult)

	errLog := new(ErrLog)
	personResult.Validate(errLog)
	// fmt.Printf("isValid: %v\n", isValid)

	fmt.Println("Сообщения:")
	for _, msg := range errLog.Messages() {
		fmt.Printf("\t%v\n", msg)
	}
}
