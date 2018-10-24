package main

import (
	"testing"
	"time"
)

var (
	USA = "USA"
)

//
// Test_Departure
//
func Test_Departure_PutShip_OutToSea(t *testing.T) {
	leo := Ship("King Leo")
	sfo := Port("San Francisco", USA)
	processor := EventProcessor()
	event := DepartureEvent(november(2, 2015), sfo, leo)

	processor.Process(event)

	if !leo.Port().Equal(AtSea()) {
		t.Errorf("port must be AtSea")
	}
}

//
// Ship
//
func Ship(name string) *ship {
	return &ship{name, AtSea()}
}

type ship struct {
	name string
	port *port
}

func (s *ship) Port() *port {
	return s.port
}

//
// AtSea()
//
func AtSea() *port {
	return &port{"AtSea", "AtSea"}
}

//
// Port
//
func Port(name, country string) *port {
	return &port{name, country}
}

type port struct {
	name    string
	country string
}

func (p *port) Equal(other *port) bool {
	return false
}

//
// EventProcessor
//
func EventProcessor() *eventProcessor {
	return &eventProcessor{}
}

type eventProcessor struct {
}

func (p *eventProcessor) Process(e *event) {

}

//
// DepartureEvent
//
func DepartureEvent(occured time.Time, port *port, ship *ship) *event {
	return &event{}
}

type event struct {
}

//
// november
//
func november(day, year int) time.Time {
	return time.Date(year, time.November, day, 0, 0, 0, 0, time.UTC)
}
