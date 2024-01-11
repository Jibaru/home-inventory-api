package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/domain/dao"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RunServer(
	host, port string,
	db *gorm.DB,
) {
	versionDAO := &dao.VersionDAO{DB: db}
	versionService := services.NewVersionService(versionDAO)
	healthController := controllers.NewHealthController(versionService)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.GET("/", healthController.Handle)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
