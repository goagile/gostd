package est

import (
	"regexp"
	"strconv"
	"fmt"
	"log"
)

type Period float64

func (p Period) String() string {
	return fmt.Sprintf("%vh", float64(p))
}

var H Period = 1.
var D Period = 8.*H
var W Period = 5.*D

var hoursRe = regexp.MustCompile(`(\d+)h`)
var daysRe = regexp.MustCompile(`(\d+)d`)
var weeksRe = regexp.MustCompile(`(\d+)w`)

func matchHours(s string) [][]string {
	return hoursRe.FindAllStringSubmatch(s, -1)
}

func matchDays(s string) [][]string {
	return daysRe.FindAllStringSubmatch(s, -1)
}

func matchWeeks(s string) [][]string {
	return weeksRe.FindAllStringSubmatch(s, -1)
}

func ParsePeriod(s string) (Period, error) {
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
