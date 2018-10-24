package main

import (
	"testing"
)

func TestGetHours(t *testing.T) {
	var source *Source = new(Source)

	var clock *DigitalClock = NewDigitalClock(source)

	source.SetTime(10, 30, 0)

	if h := clock.Hours(); h != 10 {
		t.Fatalf("%v", h)
	}
}

// --------------------------------------------------------------

type DigitalClock struct {
	source *Source
}

func NewDigitalClock(source *Source) *DigitalClock {
	c := new(DigitalClock)
	c.SetSource(source)
	return c
}

func (c *DigitalClock) SetSource(source *Source) {
	c.source = source
}

func (c *DigitalClock) Hours() int {
	return c.source.Hours()
}

// --------------------------------------------------------------

type Source struct {
	hours, minutes, seconds int
}

func (s *Source) SetTime(hours, minutes, seconds int) {
	s.hours = hours
	s.minutes = minutes
	s.seconds = seconds
}

func (s *Source) Hours() int {
	return s.hours
}
