package main

import (
	"testing"
)

func TestGetHours(t *testing.T) {
	var clock Clock
	clock = new(DigitalClock)
	clock.SetTime(10, 30, 0)

	if h := clock.Hours(); h != 10 {
		t.Fatalf("%v", h)
	}
}
