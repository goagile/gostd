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

type Clock interface {
	SetTime(h, m, s int)
	Hours() int
}

type DigitalClock struct {
	h, m, s int
}

func (c *DigitalClock) SetTime(h, m, s int) {
	c.h = h
	c.m = m
	c.s = s
}

func (c *DigitalClock) Hours() int {
	return c.h
}
