package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteBoxController struct {
	boxService *services.BoxService
}

type DeleteBoxRequest struct {
	BoxID string `param:"boxID"`
}

func NewDeleteBoxController(boxService *services.BoxService) *DeleteBoxController {
	return &DeleteBoxController{
		boxService,
	}
}

func (c *DeleteBoxController) Handle(ctx echo.Context) error {
	request := DeleteBoxRequest{}

	err := (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = c.boxService.DeleteWithTransactionsAndItemQuantities(request.BoxID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
