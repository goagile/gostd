package main


import (
	"os"
	"html/template"
)

type Person struct {
	Name string
	Age int
	Scores [3]int
}

func main() {

	t := template.Must(template.New("test").Parse(`

<h1>Person</h1>:
<p>Name: {{.Name}}</p>
<p>Age: {{.Age}}</p>
<p>Scores:</p>
<ul>
{{range .Scores}}
	<li>Score: {{.}}</li>
{{end}}
</ul>

`))

	p := Person{"Dan", 28, [3]int{1, 3, 5}}

	t.Execute(os.Stdout, p)

}
