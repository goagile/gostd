package main

import (
	"testing"
)

func Test_parse_empty(t *testing.T) {
	want := 0*H

	got, err := parse("")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
	
func Test_parse_1h(t *testing.T) {
	want := 1*H

	got, err := parse("1h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1h2h(t *testing.T) {
	want := 3*H

	got, err := parse("1h2h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1d(t *testing.T) {
	want := 1*D

	got, err := parse("1d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1d1h(t *testing.T) {
	want := 1*D+1*H

	got, err := parse("1d1h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1w(t *testing.T) {
	want := 1*W

	got, err := parse("1w")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1w2d(t *testing.T) {
	want := 1*W + 2*D

	got, err := parse("1w2d")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1w2d3h(t *testing.T) {
	want := 1*W + 2*D + 3*H

	got, err := parse("1w2d3h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_parse_1w2d3h4w5d6h(t *testing.T) {
	want := 1*W + 2*D + 3*H + 4*W + 5*D + 6*H

	got, err := parse("1w2d3h4w5d6h")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}