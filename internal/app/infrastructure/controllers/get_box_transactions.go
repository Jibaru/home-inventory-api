package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type GetBoxTransactionsController struct {
	boxService *services.BoxService
}

type GetBoxTransactionsRequest struct {
	BoxID   string `param:"boxID"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

type GetBoxTransactionsResponse struct {
	ID         string    `json:"id"`
	Type       string    `json:"type"`
	Quantity   float64   `json:"quantity"`
	BoxID      string    `json:"box_id"`
	ItemID     string    `json:"item_id"`
	ItemSku    string    `json:"item_sku"`
	ItemName   string    `json:"item_name"`
	ItemUnit   string    `json:"item_unit"`
	HappenedAt time.Time `json:"happened_at"`
}

func NewGetBoxTransactionsController(boxService *services.BoxService) *GetBoxTransactionsController {
	return &GetBoxTransactionsController{boxService}
}

func (c *GetBoxTransactionsController) Handle(ctx echo.Context) error {
	request := GetBoxTransactionsRequest{}

	err := (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindQueryParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	transactions, err := c.boxService.GetBoxTransactions(
		request.BoxID,
		services.PageFilter{
			Page: request.Page,
			Size: request.PerPage,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	total, err := c.boxService.CountBoxTransactions(request.BoxID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	responseTransactions := make([]GetBoxTransactionsResponse, len(transactions))
	for i, transaction := range transactions {
		responseTransactions[i] = GetBoxTransactionsResponse{
			ID:         transaction.ID,
			Type:       transaction.Type,
			Quantity:   transaction.Quantity,
			BoxID:      transaction.BoxID,
			ItemID:     transaction.ItemID,
			ItemSku:    transaction.ItemSku,
			ItemName:   transaction.ItemName,
			ItemUnit:   transaction.ItemUnit,
			HappenedAt: transaction.HappenedAt,
		}
	}

	return ctx.JSON(http.StatusOK, responses.NewPaginatedResponse(
		responseTransactions,
		total,
		request.Page,
		request.PerPage,
		len(transactions),
		ctx.Request().URL.Path,
	))
}
