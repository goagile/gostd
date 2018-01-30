package main

import (
	"flag"
	"fmt"
	"strings"
)


var n = flag.Bool("n", false, "пропуск символа новой строки")

func main() {
	flag.Parse()
	fmt.Println("Аргументы:")
	fmt.Print(strings.Join(flag.Args(), ", "))
	fmt.Println()
}
