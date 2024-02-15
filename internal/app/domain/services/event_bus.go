package services

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"reflect"
	"time"
)

type Event interface{}

type EventHandler func(event Event)

type EventBus interface {
	Publish(event Event) error
	Subscribe(event Event, handler EventHandler)
}

func GetEventType(event Event) string {
	return reflect.TypeOf(event).Name()
}

type StringEvent struct {
	Value string
}

type BoxItemAddedEvent struct {
	Quantity   float64
	BoxID      string
	Item       entities.Item
	HappenedAt time.Time
}

type BoxItemRemovedEvent struct {
	Quantity   float64
	BoxID      string
	Item       entities.Item
	HappenedAt time.Time
}

type ItemNotCreatedEvent struct {
	Item  entities.Item
	Asset entities.Asset
}

type ItemKeywordsNotCreatedEvent struct {
	ItemKeywords []*entities.ItemKeyword
	Asset        entities.Asset
}
