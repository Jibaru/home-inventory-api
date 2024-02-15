package memory

import "github.com/jibaru/home-inventory-api/m/internal/app/domain/services"

type EventBus struct {
	subscribers map[string][]services.EventHandler
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]services.EventHandler),
	}
}

func (e *EventBus) Publish(event services.Event) error {
	go func() {
		for _, handler := range e.subscribers[services.GetEventType(event)] {
			handler(event)
		}
	}()
	return nil
}

func (e *EventBus) Subscribe(event services.Event, handler services.EventHandler) {
	eventType := services.GetEventType(event)
	e.subscribers[eventType] = append(e.subscribers[eventType], handler)
}
