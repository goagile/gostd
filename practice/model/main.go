package main

import (
	"fmt"
	"time"
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
	msg := fmt.Sprintf("%s; %s", e.prefix, message)
	e.messages = append(e.messages, msg)
}

type Person struct {
	birth time.Time
}

// func (p Person) SetBirth(v string) error {
// 	var t time.Time
// 	t, err := time.Parse(time.RFC3339, v)
// 	if err != nil {
// 		return fmt.Errorf("ошибка формата времени RFC3339: %v ", err)
// 	}
// 	p.birth = t
// 	return nil
// }

type PersonBuilder struct {
	log    *ErrLog
	person Person
}

func (b PersonBuilder) Person() Person {
	return b.person
}

func (b PersonBuilder) BuildPerson(d map[string]interface{}) {
	b.person = Person{}

	k := "birth"
	b.log.SetPrefix(fmt.Sprintf("поле %v", k))
	if v, ok := d[k]; ok {
		b.setBirth(v)
	} else {
		b.log.Add(fmt.Sprintf("не найдено"))
	}
}

func (b PersonBuilder) setBirth(v interface{}) {
	var birth string
	birth, ok := v.(string)
	if !ok {
		b.log.Add(fmt.Sprintf("неверный тип, ожидается string "))
		return
	}
	var t time.Time
	t, err := time.Parse(time.RFC3339, birth)
	if err != nil {
		b.log.Add(fmt.Sprintf("ошибка формата времени RFC3339; %v", err))
		return
	}
	b.person.birth = t
}

func main() {
	d := map[string]interface{}{
		"birth": "2006-01-02T15:04:05Z",
		// "birth": "`{sdsd}`",
		// "birth": 34,
	}
	log := &ErrLog{}

	fmt.Println("Person:")
	b := PersonBuilder{log: log}
	b.BuildPerson(d)
	fmt.Printf("\t%+v\n", b.Person())

	fmt.Println("Messages:")
	for _, m := range log.Messages() {
		fmt.Printf("\t%v\n", m)
	}
}
