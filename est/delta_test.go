package est

import (
	"testing"
	"fmt"
)


//
// ParseDelta
//
func Test_ParseDelta_0percents(t *testing.T) {
	want := Delta(0.)

	got, err := ParseDelta("0%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParseDelta_10percents(t *testing.T) {
	want := Delta(0.1)

	got, err := ParseDelta("10%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParseDelta_100percents(t *testing.T) {
	want := Delta(1.)

	got, err := ParseDelta("100%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParseDelta_120percents(t *testing.T) {
	want := Delta(1.2)

	got, err := ParseDelta("120%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_ParseDelta_point5percents(t *testing.T) {
	want := Delta(0.05)

	got, err := ParseDelta(".5%")
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
func Test_String_40percent(t *testing.T) {
	want := "40%"

	d, err := ParseDelta("40%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := fmt.Sprintf("%v", d)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Int
//
func Test_Int(t *testing.T) {
	want := 65

	d, err := ParseDelta("65%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := d.Int()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Float64
//
func Test_Float64(t *testing.T) {
	want := 0.4

	d, err := ParseDelta("40%")
	if err != nil {
		t.Fatalf("\nerr:%v\n", err)
	}

	got := d.Float64()

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}
