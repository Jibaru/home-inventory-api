package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type CreateItemController struct {
	itemService *services.ItemService
}

type CreateItemRequest struct {
	Sku         string   `form:"sku"`
	Name        string   `form:"name"`
	Description *string  `form:"description"`
	Unit        string   `form:"unit"`
	Keywords    []string `form:"keywords[]"`
}

type CreateItemResponse struct {
	ID          string   `json:"id"`
	Sku         string   `json:"sku"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Unit        string   `json:"unit"`
	Keywords    []string `json:"keywords"`
}

func NewCreateItemController(
	itemService *services.ItemService,
) *CreateItemController {
	return &CreateItemController{
		itemService,
	}
}

func (c *CreateItemController) Handle(ctx echo.Context) error {
	userID := ctx.Get("auth_id").(string)
	request := CreateItemRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
	}
	tempDir, tempFile, err := mapFileHeaderToTempFolderAndFile(fileHeader)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
	}
	defer os.RemoveAll(tempDir)
	defer tempFile.Close()

	item, err := c.itemService.Create(
		request.Sku,
		request.Name,
		request.Description,
		request.Unit,
		userID,
		request.Keywords,
		tempFile,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.NewDataResponse(&CreateItemResponse{
		ID:          item.ID,
		Sku:         item.Sku,
		Name:        item.Name,
		Description: item.Description,
		Unit:        item.Unit,
		Keywords:    request.Keywords,
	}))
}
