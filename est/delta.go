package est

import (
	"regexp"
	"strconv"
	"fmt"
	"log"
)

type Delta float64

func (d Delta) Int() int {
	return int(float64(d) * 100.)
}

func (d Delta) Float64() float64 {
	return float64(d)
}

func (d Delta) String() string {
	return fmt.Sprintf("%v%%", d.Int())
}

var deltaRe = regexp.MustCompile(`(\d+)%`)

func ParseDelta(s string) (Delta, error) {
	m := deltaRe.FindStringSubmatch(s)
	return delta(m) / 100, nil
}

func delta(m []string) Delta {
	if len(m) <= 1 {
		return Delta(0)
	}
	
	v := m[1]
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Printf("delta ParseFloat %v\n", err)
		return Delta(0)
	}

	return Delta(f)
}
