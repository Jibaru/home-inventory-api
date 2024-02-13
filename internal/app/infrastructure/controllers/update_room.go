package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateRoomController struct {
	roomService *services.RoomService
}

type UpdateRoomRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	RoomID      string  `param:"roomID"`
}

type UpdateRoomResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewUpdateRoomController(roomService *services.RoomService) *UpdateRoomController {
	return &UpdateRoomController{
		roomService,
	}
}

func (c *UpdateRoomController) Handle(ctx echo.Context) error {
	request := UpdateRoomRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	room, err := c.roomService.Update(request.RoomID, request.Name, request.Description)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.NewDataResponse(
		&UpdateRoomResponse{
			ID:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		},
	))
}
