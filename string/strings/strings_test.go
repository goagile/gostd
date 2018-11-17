package strings

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

//
// Split
//
func Test_Split(t *testing.T) {
	want := []string{"a", "b", "c"}

	got := strings.Split("a.b.c", ".")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
	}
}

//
// Join
//
func Test_Join(t *testing.T) {
	want := "a.b.c"

	got := strings.Join([]string{"a", "b", "c"}, ".")

	if want != got {
		t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
	}
}

//
// Count
//
func Test_Count(t *testing.T) {
	tests := []struct {
		str  string
		sub  string
		want int
	}{
		{"a", "a", 1},
		{"bb", "b", 2},
		{"abbcccdcc", "c", 5},
	}

	for _, test := range tests {

		got := strings.Count(test.str, test.sub)

		want := test.want
		if want != got {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}

//
// Contains
//
func Test_Contains(t *testing.T) {
	tests := []struct {
		str  string
		sub  string
		want bool
	}{
		{"a", "a", true},
		{"bb", "a", false},
		{"abbcccdcc", "c", true},
	}

	for _, test := range tests {

		got := strings.Contains(test.str, test.sub)

		want := test.want
		if want != got {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}

//
// HasPrefix
//
func Test_HasPrefix(t *testing.T) {
	tests := []struct {
		str  string
		sub  string
		want bool
	}{
		{"TEST_XXX12131", "TEST", true},
		{"PROD_123114Fs", "TEST", false},
	}

	for _, test := range tests {

		got := strings.HasPrefix(test.str, test.sub)

		want := test.want
		if want != got {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}

//
// HasSuffix
//
func Test_HasSuffix(t *testing.T) {
	tests := []struct {
		str  string
		sub  string
		want bool
	}{
		{"XXX12131_TEST", "TEST", true},
		{"123114Fs_PROD", "TEST", false},
	}

	for _, test := range tests {

		got := strings.HasSuffix(test.str, test.sub)

		want := test.want
		if want != got {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}

//
// Map
//
func Test_Map(t *testing.T) {
	want := "11111"

	s := "01010"

	zeroToOne := func(r rune) rune {
		if r == '0' {
			return '1'
		}
		return r
	}

	got := strings.Map(zeroToOne, s)

	if want != got {
		t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
	}
}

//
// Fields
//
func Test_Fields(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{" aa bb cc ", []string{"aa", "bb", "cc"}},
	}

	for _, test := range tests {

		got := strings.Fields(test.str)

		want := test.want
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}

//
// FieldsFunc
//
func Test_FieldsFunc(t *testing.T) {
	tests := []struct {
		str  string
		want []string
	}{
		{" aa 11 22 bb ", []string{"11", "22"}},
	}

	isNotDigit := func(r rune) bool {
		return !unicode.IsDigit(r)
	}

	for _, test := range tests {

		got := strings.FieldsFunc(test.str, isNotDigit)

		want := test.want
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("\nwant: %v\ngot: %v\n", want, got)
		}
	}
}
