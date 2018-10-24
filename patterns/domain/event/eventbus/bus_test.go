package main

import "testing"

func Test_Bus_Raise(t *testing.T) {

	// create Event handler
	handler := EmailNotificationHandler()

	// create Event Bus
	bus := Bus()

	// register Event handler to handle Event on Bus
	bus.Register(handler)

	// raise domain event
	bus.Raise(UserAchieveTopLevel())

	if !handler.wasRaised {
		t.Errorf("Handle must be Raised")
	}
}

//
// Bus
//
func Bus() *bus {
	return &bus{}
}

type bus struct {
	handlers []Handler
}

func (b *bus) Register(handler Handler) {
	b.handlers = append(b.handlers, handler)
}

func (b *bus) Raise(event *event) {
	for _, h := range b.handlers {
		h.Handle(event)
	}
}

//
// UserAchieveTopLevelHandler
//
func UserAchieveTopLevel() *event {
	return &event{}
}

type event struct{}

//
// Handler
//
type Handler interface {
	Handle(e *event)
}

//
// EmailNotificationHandler
//
func EmailNotificationHandler() *emailNotificationHandler {
	return &emailNotificationHandler{wasRaised: false}
}

type emailNotificationHandler struct {
	wasRaised bool
}

func (h *emailNotificationHandler) Handle(e *event) {
	h.wasRaised = true
}
