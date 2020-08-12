package main

import (
	"regexp"
	"strconv"
	// "fmt"
	"log"
)

func main() {

}

type Period float64

var H Period = 1.
var D Period = 8.*H
var W Period = 5.*D

var HoursRe = regexp.MustCompile(`(\d+)h`)
var DaysRe = regexp.MustCompile(`(\d+)d`)
var WeeksRe = regexp.MustCompile(`(\d+)w`)

func matchHours(s string) [][]string {
	return HoursRe.FindAllStringSubmatch(s, -1)
}

func matchDays(s string) [][]string {
	return DaysRe.FindAllStringSubmatch(s, -1)
}

func matchWeeks(s string) [][]string {
	return WeeksRe.FindAllStringSubmatch(s, -1)
}

func parse(s string) (Period, error) {
	h := 0 * H

	for _, m := range matchHours(s) {
		h += period(m) * H
	}

	for _, m := range matchDays(s) {
		h += period(m) * D
	}

	for _, m := range matchWeeks(s) {
		h += period(m) * W
	}

	return h, nil
}

func period(m []string) Period {
	if len(m) <= 1 {
		return Period(0)
	}
	
	v := m[1]
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("findPeriod ParseFloat %v\n", err)
		return Period(0)
	}

	return Period(f)
}
