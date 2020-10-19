package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("SELECT ")
	b.WriteString("id,first,last ")
	b.WriteString("FROM ")
	b.WriteString("clients")
	fmt.Println(b.String())
}
