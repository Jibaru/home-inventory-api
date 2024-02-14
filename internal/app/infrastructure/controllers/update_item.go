package controllers

import (
	"errors"
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type UpdateItemController struct {
	itemService *services.ItemService
}

type UpdateItemRequest struct {
	ItemID      string   `param:"itemID"`
	Sku         string   `form:"sku"`
	Name        string   `form:"name"`
	Description *string  `form:"description"`
	Unit        string   `form:"unit"`
	Keywords    []string `form:"keywords[]"`
}

type UpdateItemResponse struct {
	ID          string   `json:"id"`
	Sku         string   `json:"sku"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Unit        string   `json:"unit"`
	Keywords    []string `json:"keywords"`
}

func NewUpdateItemController(itemService *services.ItemService) *UpdateItemController {
	return &UpdateItemController{
		itemService: itemService,
	}
}

func (c *UpdateItemController) Handle(ctx echo.Context) error {
	request := UpdateItemRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}
	err = (&echo.DefaultBinder{}).BindPathParams(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	fileHeader, err := ctx.FormFile("file")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		logger.LogError(err)
		return ctx.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
	}

	var tempFile *os.File
	if fileHeader != nil {
		var tempDir string
		tempDir, tempFile, err = mapFileHeaderToTempFolderAndFile(fileHeader)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		defer os.RemoveAll(tempDir)
		defer tempFile.Close()
	}

	item, err := c.itemService.Update(
		request.ItemID,
		request.Sku,
		request.Name,
		request.Description,
		request.Unit,
		request.Keywords,
		tempFile,
	)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	keywords := make([]string, len(item.Keywords))
	for i, keyword := range item.Keywords {
		keywords[i] = keyword.Value
	}

	return ctx.JSON(http.StatusOK, responses.NewDataResponse(&UpdateItemResponse{
		ID:          item.ID,
		Sku:         item.Sku,
		Name:        item.Name,
		Description: item.Description,
		Unit:        item.Unit,
		Keywords:    keywords,
	}))
}
