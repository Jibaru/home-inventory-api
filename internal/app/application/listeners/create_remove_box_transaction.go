package listeners

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	domain "github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/jibaru/home-inventory-api/m/logger"
)

type CreateRemoveBoxTransactionListener struct {
	boxService *services.BoxService
}

func NewCreateRemoveBoxTransactionListener(
	boxService *services.BoxService,
) *CreateRemoveBoxTransactionListener {
	return &CreateRemoveBoxTransactionListener{
		boxService: boxService,
	}
}

func (l *CreateRemoveBoxTransactionListener) Handle(event domain.Event) {
	if event, ok := event.(domain.BoxItemRemovedEvent); ok {
		_, err := l.boxService.CreateRemoveBoxTransaction(
			event.Quantity,
			event.BoxID,
			event.Item,
			event.HappenedAt,
		)
		if err != nil {
			logger.LogError(err)
		}
	}
}
