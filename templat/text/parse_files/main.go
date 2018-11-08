package main

import (
	"bytes"
	"log"
	"text/template"
)

type Email struct {
	From    string
	To      string
	Subject string
}

func main() {
	// парсим файл в шаблон
	t := template.Must(template.ParseFiles("tpl.txt"))

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
