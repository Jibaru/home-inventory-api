package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RemoveItemFromBoxController struct {
	boxService *services.BoxService
}

type RemoveItemFromBoxRequest struct {
	Quantity float64 `json:"quantity"`
	ItemID   string  `param:"itemID"`
	BoxID    string  `param:"boxID"`
}

func NewRemoveItemFromBoxController(
	boxService *services.BoxService,
) *RemoveItemFromBoxController {
	return &RemoveItemFromBoxController{boxService}
}

func (c *RemoveItemFromBoxController) Handle(ctx echo.Context) error {
	request := RemoveItemFromBoxRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = c.boxService.RemoveItemFromBox(
		request.Quantity,
		request.BoxID,
		request.ItemID,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
