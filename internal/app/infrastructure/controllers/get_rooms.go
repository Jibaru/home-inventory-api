package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetRoomsController struct {
	roomService *services.RoomService
}

type GetRoomsRequest struct {
	Search  string `query:"search"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

type GetRoomsResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func NewGetRoomsController(
	roomService *services.RoomService,
) *GetRoomsController {
	return &GetRoomsController{roomService: roomService}
}

func (c *GetRoomsController) Handle(ctx echo.Context) error {
	userID := ctx.Get("auth_id").(string)
	request := GetRoomsRequest{}

	err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	rooms, err := c.roomService.GetAll(
		request.Search,
		userID,
		services.PageFilter{
			Page: request.Page,
			Size: request.PerPage,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	total, err := c.roomService.CountAll(
		request.Search,
		userID,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	responseRooms := make([]*GetRoomsResponse, 0)
	for _, room := range rooms {
		responseRooms = append(responseRooms, &GetRoomsResponse{
			ID:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		})
	}

	return ctx.JSON(http.StatusOK, responses.NewPaginatedResponse(
		responseRooms,
		total,
		request.Page,
		request.PerPage,
		len(rooms),
		ctx.Request().URL.Path,
	))
}
