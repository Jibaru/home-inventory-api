package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChangeBoxRoomController struct {
	boxService *services.BoxService
}

type ChangeBoxRoomRequest struct {
	BoxID  string `param:"boxID"`
	RoomID string `json:"room_id"`
}

func NewChangeBoxRoomController(boxService *services.BoxService) *ChangeBoxRoomController {
	return &ChangeBoxRoomController{
		boxService: boxService,
	}
}

func (c *ChangeBoxRoomController) Handle(ctx echo.Context) error {
	request := ChangeBoxRoomRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = c.boxService.TransferToRoom(request.BoxID, request.RoomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
