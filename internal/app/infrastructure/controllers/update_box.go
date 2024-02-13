package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateBoxController struct {
	boxService *services.BoxService
}

type UpdateBoxRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	BoxID       string  `param:"boxID"`
}

type UpdateBoxResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewUpdateBoxController(boxService *services.BoxService) *UpdateBoxController {
	return &UpdateBoxController{
		boxService,
	}
}

func (c *UpdateBoxController) Handle(ctx echo.Context) error {
	request := UpdateBoxRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	box, err := c.boxService.Update(request.BoxID, request.Name, request.Description)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.NewDataResponse(
		&UpdateBoxResponse{
			ID:          box.ID,
			Name:        box.Name,
			Description: box.Description,
		},
	))
}
