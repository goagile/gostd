package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', 0)
	format := "|\t%v\t|\t%v\t|\t%v\t|\n"

	fmt.Fprintf(w, format, "A", "B", "C")
	fmt.Fprintf(w, format, "--", "--", "--")
	fmt.Fprintf(w, format, "sfsdfsfds", "fsddf", "dd")
	for i := 0; i < 5; i++ {
		fmt.Fprintf(w, format, i, i, i)
	}
	fmt.Fprintf(w, format, "ss", "rerwrewervdsgds", "z")

	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}
