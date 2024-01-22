package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/entities"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

type CreateAssetController struct {
	assetService *services.AssetService
}

type CreateAssetResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Extension  string `json:"extension"`
	Size       int64  `json:"size"`
	FileID     string `json:"file_id"`
	EntityID   string `json:"entity_id"`
	EntityName string `json:"entity_name"`
	Url        string `json:"url"`
}

func NewCreateAssetController(
	assetService *services.AssetService,
) *CreateAssetController {
	return &CreateAssetController{assetService}
}

func (c *CreateAssetController) Handle(ctx echo.Context) error {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return err
	}
	tempDir, tempFile, err := mapFileHeaderToTempFolderAndFile(fileHeader)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
	}
	defer os.RemoveAll(tempDir)
	defer tempFile.Close()

	userID := ctx.Get("auth_id").(string)

	asset, err := c.assetService.CreateFromFile(tempFile, entities.NewIdentifiableEntity(userID))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.NewDataResponse(&CreateAssetResponse{
		ID:         asset.ID,
		Name:       asset.Name,
		Extension:  asset.Extension,
		Size:       asset.Size,
		FileID:     asset.FileID,
		EntityID:   asset.EntityID,
		EntityName: asset.EntityName,
		Url:        c.assetService.GetUrl(asset),
	}))
}
