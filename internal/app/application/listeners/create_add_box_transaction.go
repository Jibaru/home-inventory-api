package listeners

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	domain "github.com/jibaru/home-inventory-api/m/internal/app/domain/services"
	"github.com/jibaru/home-inventory-api/m/logger"
)

type CreateAddBoxTransactionListener struct {
	boxService *services.BoxService
}

func NewCreateAddBoxTransactionListener(
	boxService *services.BoxService,
) *CreateAddBoxTransactionListener {
	return &CreateAddBoxTransactionListener{
		boxService: boxService,
	}
}

func (l *CreateAddBoxTransactionListener) Handle(event domain.Event) {
	if e, ok := event.(domain.BoxItemAddedEvent); ok {
		_, err := l.boxService.CreateAddBoxTransaction(
			e.Quantity,
			e.BoxID,
			e.Item,
			e.HappenedAt,
		)
		if err != nil {
			logger.LogError(err)
		}
	}
}
