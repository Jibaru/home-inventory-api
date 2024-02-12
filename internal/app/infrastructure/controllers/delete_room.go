package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteRoomController struct {
	roomService *services.RoomService
}

type DeleteRoomRequest struct {
	RoomID string `param:"roomID"`
}

func NewDeleteRoomController(roomService *services.RoomService) *DeleteRoomController {
	return &DeleteRoomController{
		roomService,
	}
}

func (c *DeleteRoomController) Handle(ctx echo.Context) error {
	request := DeleteRoomRequest{}

	err := (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = c.roomService.Delete(request.RoomID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
