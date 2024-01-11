package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	ApiVersion = "0.0.1"
)

func RunServer(host, port string) {
	healthController := controllers.NewHealthController(ApiVersion)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.GET("/", healthController.Handle)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
