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
	userDAO := dao.UserDAO{DB: db}

	userService := services.NewUserService(userDAO)
	versionService := services.NewVersionService(versionDAO)

	healthController := controllers.NewHealthController(versionService)
	signOnController := controllers.NewSignOnController(userService)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.GET("/", healthController.Handle)
	api.POST("/users", signOnController.Handle)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
