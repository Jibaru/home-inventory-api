package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthController struct {
	versionService *services.VersionService
}

type HealthControllerResponse struct {
	Version string `json:"version"`
}

func NewHealthController(versionService *services.VersionService) *HealthController {
	return &HealthController{versionService}
}

func (c *HealthController) Handle(ctx echo.Context) error {
	version, err := c.versionService.GetLatestVersion()
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			responses.MessageResponse{Message: err.Error()},
		)
	}

	return ctx.JSON(
		http.StatusOK,
		HealthControllerResponse{version.Value},
	)
}
