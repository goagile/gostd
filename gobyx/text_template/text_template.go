package main


import (
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age int
	Scores [3]int
}

func main() {

	t := template.Must(template.New("test").Parse(`

Person:
  Name: {{.Name}}
  Age: {{.Age}}
  Scores:{{range .Scores}} 
  	Score: {{.}}{{end}}

`))

	p := Person{"Dan", 28, [3]int{1, 3, 5}}

	t.Execute(os.Stdout, p)

}
