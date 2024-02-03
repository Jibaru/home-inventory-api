package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetBoxesController struct {
	boxService *services.BoxService
}

type GetBoxesRequest struct {
	RoomID  string `query:"room_id"`
	Search  string `query:"search"`
	Page    int    `query:"page"`
	PerPage int    `query:"per_page"`
}

type GetBoxesResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	RoomID      string  `json:"room_id"`
}

func NewGetBoxesController(boxService *services.BoxService) *GetBoxesController {
	return &GetBoxesController{boxService}
}

func (c *GetBoxesController) Handle(ctx echo.Context) error {
	userID := ctx.Get("auth_id").(string)
	request := GetBoxesRequest{}

	err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	boxes, err := c.boxService.GetAll(
		request.RoomID,
		userID,
		request.Search,
		services.PageFilter{
			Page: request.Page,
			Size: request.PerPage,
		},
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	total, err := c.boxService.CountAll(
		userID,
		request.Search,
		request.RoomID,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	responseBoxes := make([]*GetBoxesResponse, 0)
	for _, box := range boxes {
		responseBoxes = append(responseBoxes, &GetBoxesResponse{
			ID:          box.ID,
			Name:        box.Name,
			Description: box.Description,
			RoomID:      box.RoomID,
		})
	}

	return ctx.JSON(http.StatusOK, responses.NewPaginatedResponse(
		responseBoxes,
		total,
		request.Page,
		request.PerPage,
		len(boxes),
		ctx.Request().URL.Path,
	))
}
