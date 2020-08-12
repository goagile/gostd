package main

import (
	"fmt"
	"log"
	"flag"
	"github.com/khardi/gostd/est"
)

var period string
var delta string

func init() {
	flag.StringVar(&period, "period", "0h", "Period")
	flag.StringVar(&delta, "delta", "0%", "Delta")
	flag.Parse()
}

func main() {
	p, err := est.ParsePeriod(period)
	if err != nil {
		log.Fatalf("Parse period:%v err:%v\n", period, err)
	}

	d, err := est.ParseDelta(delta)
	if err != nil {
		log.Fatalf("Parse delta:%v err:%v\n", delta, err)
	}

	add := est.Period(float64(p) + float64(p) * float64(d))
	fmt.Printf("%v + %v = %v\n", p, d, add)

	// fmt.Printf("period: %v\n", p)
	// fmt.Printf("delta: %v\n", d)
	// fmt.Printf("period + delta: %v\n", est.Period(float64(p) + float64(p) * float64(d)))
}
