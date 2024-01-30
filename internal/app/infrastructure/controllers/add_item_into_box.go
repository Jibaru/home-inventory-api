package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AddItemIntoBoxController struct {
	boxService *services.BoxService
}

type AddItemIntoBoxRequest struct {
	Quantity float64 `json:"quantity"`
	ItemID   string  `json:"item_id"`
	BoxID    string  `param:"boxID"`
}

type AddItemIntoBoxResponse struct {
	ID       string  `json:"id"`
	Quantity float64 `json:"quantity"`
	ItemID   string  `json:"item_id"`
	BoxID    string  `json:"box_id"`
}

func NewAddItemIntoBoxController(boxService *services.BoxService) *AddItemIntoBoxController {
	return &AddItemIntoBoxController{
		boxService,
	}
}

func (c *AddItemIntoBoxController) Handle(ctx echo.Context) error {
	request := AddItemIntoBoxRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	boxItem, err := c.boxService.AddItemIntoBox(
		request.Quantity,
		request.BoxID,
		request.ItemID,
	)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, responses.NewDataResponse(&AddItemIntoBoxResponse{
		ID:       boxItem.ID,
		Quantity: boxItem.Quantity,
		ItemID:   boxItem.ItemID,
		BoxID:    boxItem.BoxID,
	}))
}
