package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransferItemController struct {
	boxService *services.BoxService
}

type TransferItemRequest struct {
	BoxID            string `param:"boxID"`
	ItemID           string `param:"itemID"`
	BoxDestinationID string `json:"box_destination_id"`
}

func NewTransferItemController(boxService *services.BoxService) *TransferItemController {
	return &TransferItemController{boxService}
}

func (c *TransferItemController) Handle(ctx echo.Context) error {
	request := TransferItemRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = c.boxService.TransferItem(
		request.BoxID,
		request.BoxDestinationID,
		request.ItemID,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
