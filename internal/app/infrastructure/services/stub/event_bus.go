package stub

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/stretchr/testify/mock"
)

type EventBusMock struct {
	mock.Mock
}

func (m *EventBusMock) Publish(event services.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *EventBusMock) Subscribe(event services.Event, handler services.EventHandler) {
	m.Called(event, handler)
}
