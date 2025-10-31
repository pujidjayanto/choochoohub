package eventbus

import (
	evbus "github.com/asaskevich/EventBus"
)

type EventBus interface {
	Publish(event string, payload interface{})
	Subscribe(event string, handler func(payload interface{}))
}

type eventBus struct {
	bus evbus.Bus
}

func New() EventBus {
	return &eventBus{
		bus: evbus.New(),
	}
}

func (e *eventBus) Publish(event string, payload interface{}) {
	e.bus.Publish(event, payload, false)
}

func (e *eventBus) Subscribe(event string, handler func(payload interface{})) {
	e.bus.Subscribe(event, handler)
}
