package main

import (
	"flag"
	"fmt"
)

func main() {
	var d Duration
	flag.CommandLine.Var(&d, "d", "duration")

	flag.Parse()

	fmt.Printf("%v\n", d.String())
}

type Duration struct {
	v int
	u string
}

func (d *Duration) String() string {
	return fmt.Sprintf("%v %s", d.v, d.u)
}

func (d *Duration) Set(s string) error {
	var unit string
	fmt.Sscanf(s, "%v%v", &d.v, &unit)
	switch unit {
	case "d":
		d.u = "day"
		return nil
	}
	return fmt.Errorf("Unknown unit: %v", unit)
}
