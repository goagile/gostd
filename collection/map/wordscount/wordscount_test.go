package main

import (
	// "fmt"
	"testing"
	"reflect"
	"strings"
)

type result map[string]int

func TestA(t *testing.T) {
	want := result{
		"Go": 2,
		"to": 1,
		"learn": 1,
	}

	got := count("Go to learn Go")

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\ngot:%v\nwant:%v\n", got, want)
	}
}

func count(s string) result {
	r := make(result)
	for _, w := range strings.Split(s, " ") {
		r[w]++
	}
	return r
}
