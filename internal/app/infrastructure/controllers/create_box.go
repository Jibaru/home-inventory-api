package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateBoxController struct {
	boxService *services.BoxService
}

type CreateBoxRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	RoomID      string  `param:"roomID"`
}

type CreateBoxResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewCreateBoxController(boxService *services.BoxService) *CreateBoxController {
	return &CreateBoxController{boxService}
}

func (c *CreateBoxController) Handle(ctx echo.Context) error {
	request := CreateBoxRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	room, err := c.boxService.Create(request.Name, request.Description, request.RoomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.NewDataResponse(
		&CreateBoxResponse{
			ID:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		},
	))
}
