package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	repositories "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func RunServer(
	host, port string,
	db *gorm.DB,
) {
	versionRepository := repositories.NewVersionRepository(db)
	userRepository := repositories.NewUserRepository(db)

	userService := services.NewUserService(userRepository)
	versionService := services.NewVersionService(versionRepository)

	healthController := controllers.NewHealthController(versionService)
	signOnController := controllers.NewSignOnController(userService)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.GET("/", healthController.Handle)
	api.POST("/users", signOnController.Handle)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
