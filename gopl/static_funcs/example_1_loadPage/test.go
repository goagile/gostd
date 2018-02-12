package main

import (
	"fmt"
	"io/ioutil"
)


type Page struct {
	title string
	body []byte
}


func (p *Page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}


func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title: title, body: body}, nil
}


func main() {

	p1 := &Page{title: "TetsPage", body: []byte("This is a simple Page.")}
	p1.save()
	p2, _ := loadPage("TetsPage")
	fmt.Println(string(p2.body))

}
