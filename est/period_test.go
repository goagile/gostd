package est

import (
	"testing"
	"fmt"
)


//
// ParsePeriod
//
func Test_ParsePeriod_empty(t *testing.T) {
	want := 0*H

	got, err := ParsePeriod("")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
	
func Test_ParsePeriod_1h(t *testing.T) {
	want := 1*H

	got, err := ParsePeriod("1h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1h2h(t *testing.T) {
	want := 3*H

	got, err := ParsePeriod("1h2h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1d(t *testing.T) {
	want := 1*D

	got, err := ParsePeriod("1d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1d1h(t *testing.T) {
	want := 1*D+1*H

	got, err := ParsePeriod("1d1h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1w(t *testing.T) {
	want := 1*W

	got, err := ParsePeriod("1w")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1w2d(t *testing.T) {
	want := 1*W + 2*D

	got, err := ParsePeriod("1w2d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1w2d3h(t *testing.T) {
	want := 1*W + 2*D + 3*H

	got, err := ParsePeriod("1w2d3h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParsePeriod_1w2d3h4w5d6h(t *testing.T) {
	want := 1*W + 2*D + 3*H + 4*W + 5*D + 6*H

	got, err := ParsePeriod("1w2d3h4w5d6h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}


//
// String
//
func Test_String_9h(t *testing.T) {
	want := "9h"

	p, err := ParsePeriod("9h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := fmt.Sprintf("%v", p)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_String_1d(t *testing.T) {
	want := "8h"

	p, err := ParsePeriod("1d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := fmt.Sprintf("%v", p)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_String_1w1d(t *testing.T) {
	want := "48h"

	p, err := ParsePeriod("1w1d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := fmt.Sprintf("%v", p)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
