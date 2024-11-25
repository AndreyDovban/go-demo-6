package event

const EventLinkVisited = "LinkVisited"

type Event struct {
	Type string
	Data any
}

type EventBus struct {
	bus chan Event
}

func NewEventBus() *EventBus {
	return &EventBus{
		bus: make(chan Event),
	}
}

func (e *EventBus) Publish(event Event) {
	e.bus <- event
}

func (e *EventBus) Subscribe() <-chan Event {
	return e.bus
}
