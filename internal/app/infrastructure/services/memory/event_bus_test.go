package memory

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewEventBus(t *testing.T) {
	eventBus := NewEventBus()

	assert.NotNil(t, eventBus)
	assert.NotNil(t, eventBus.subscribers)
	assert.Len(t, eventBus.subscribers, 0)
}

func TestEventBusSubscribeAndPublish(t *testing.T) {
	calledValue := ""
	subscriber := func(event services.Event) {
		calledValue = event.(services.StringEvent).Value
	}

	eventBus := NewEventBus()
	eventBus.Subscribe(services.StringEvent{}, subscriber)

	err := eventBus.Publish(services.StringEvent{
		Value: "hello",
	})

	time.Sleep(1 * time.Second)

	assert.Nil(t, err)
	assert.Equal(t, "hello", calledValue)
}
