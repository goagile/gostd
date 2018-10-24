package main

import (
	"testing"
)

func TestGetHours(t *testing.T) {
	var clock *DigitalClock = new(DigitalClock)

	var source *Source = new(Source)
	source.RegisterClock(clock)

	source.SetTime(10, 30, 0)

	if h := clock.Hours(); h != 10 {
		t.Fatalf("%v", h)
	}
}

// --------------------------------------------------------------

type DigitalClock struct {
	hours, minutes, seconds int
}

func (c *DigitalClock) Hours() int {
	return c.hours
}

func (c *DigitalClock) Update(hours, minutes, seconds int) {
	c.hours = hours
	c.minutes = minutes
	c.seconds = seconds
}

// --------------------------------------------------------------

type Source struct {
	hours, minutes, seconds int
	clocks                  []*DigitalClock
}

func (s *Source) SetTime(hours, minutes, seconds int) {
	s.hours = hours
	s.minutes = minutes
	s.seconds = seconds
	s.notifyClocks()
}

func (s *Source) RegisterClock(clock *DigitalClock) {
	s.clocks = append(s.clocks, clock)
}

func (s *Source) notifyClocks() {
	for _, c := range s.clocks {
		c.Update(s.hours, s.minutes, s.seconds)
	}
}
