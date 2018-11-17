package main

import (
	"bytes"
	"log"
	"text/template"
)

var tpl = `

	New Email:
		Subject: {{.Subject}}
		From: {{.From}}
		To: {{.To}}

`

type Email struct {
	From    string
	To      string
	Subject string
}

func main() {
	// создаем новый шаблон и даем ему имя
	t := template.New("data")

	// парсим строку в синтаксическое дерево
	t, err := t.Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	// создаем контекст шаблона
	e := &Email{
		Subject: "Hello!",
		From:    "aaa@gmail.com",
		To:      "bbb@yandex.ru",
	}

	// создаем буффер для хранения результатов
	b := new(bytes.Buffer)

	// подставляем контекст в шаблон и сохраняем в буфер
	t.Execute(b, e)

	log.Println(b.String())
}
