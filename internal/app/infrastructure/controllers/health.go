package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthController struct {
	Version string
}

type HealthControllerResponse struct {
	Version string `json:"version"`
}

func NewHealthController(version string) *HealthController {
	return &HealthController{
		version,
	}
}

func (c *HealthController) Handle(ctx echo.Context) error {
	return ctx.JSON(
		http.StatusOK,
		HealthControllerResponse{c.Version},
	)
}
