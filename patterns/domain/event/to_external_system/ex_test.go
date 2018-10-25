package to_external_system

import (
	"testing"
	"time"
)

var (
	USA      = "USA"
	Registry = ServiceRegistry(ExternalSystemGateway())
)

func Test_HandleArrival_And_Send_Message_To_External_System(t *testing.T) {
	sfo := Port("San Francisco", USA)

	sfo.HandleArrival(ArrivalEvent(october(20, 2008), sfo))

	if !Registry.ExternalSystemGateway().IsReceiveMessage() {
		t.Errorf("External System has not receive message")
	}
}

func october(day, year int) time.Time {
	return time.Date(year, time.October, day, 0, 0, 0, 0, time.UTC)
}

//
// ArrivalEvent
//
func ArrivalEvent(occured time.Time, port *port) *event {
	return &event{occured, port}
}

type event struct {
	occured time.Time
	port    *port
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

func (p *port) HandleArrival(e *event) {
	// ...
	Registry.ExternalSystemGateway().Notify(e.occured, e.port.name, e.port.country)
	// ...
}

//
// ServiceRegistry
//
func ServiceRegistry(gateway *gateway) *serviceRegistry {
	return &serviceRegistry{gateway}
}

type serviceRegistry struct {
	gateway *gateway
}

func (r *serviceRegistry) ExternalSystemGateway() *gateway {
	return r.gateway
}

//
// ExternalSystemGateway
//
func ExternalSystemGateway() *gateway {
	isReceiveMessage := false
	return &gateway{isReceiveMessage}
}

type gateway struct {
	isReceiveMessage bool
}

func (g *gateway) IsReceiveMessage() bool {
	return g.isReceiveMessage
}

func (g *gateway) Notify(occured time.Time, name, country string) {
	g.Send(BuildArrivalMessage(occured, name, country))
}

func (g *gateway) Send(message map[string]interface{}) {
	// ...
	g.isReceiveMessage = true
	// ...
}

//
// BuildArrivalMessage
//
func BuildArrivalMessage(occured time.Time, name, country string) map[string]interface{} {
	return map[string]interface{}{
		"occured":      occured.Format(time.RFC3339),
		"port_name":    name,
		"port_country": country,
	}
}
