package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateRoomController struct {
	roomService *services.RoomService
}

type CreateRoomRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateRoomResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewCreateRoomController(
	roomService *services.RoomService,
) *CreateRoomController {
	return &CreateRoomController{
		roomService,
	}
}

func (c *CreateRoomController) Handle(ctx echo.Context) error {
	request := CreateRoomRequest{}
	userID := ctx.Get("auth_id").(string)

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	room, err := c.roomService.Create(request.Name, request.Description, userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.NewDataResponse(
		&CreateRoomResponse{
			ID:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		},
	))
}
