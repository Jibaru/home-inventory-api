package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/auth/jwt"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	repositories "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"time"
)

func RunServer(
	host string,
	port string,
	jwtSecret string,
	jwtDuration time.Duration,
	db *gorm.DB,
) {
	tokenGenerator := jwt.NewJwtGenerator(jwtSecret, jwtDuration)

	versionRepository := repositories.NewVersionRepository(db)
	userRepository := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepository, tokenGenerator)
	userService := services.NewUserService(userRepository)
	versionService := services.NewVersionService(versionRepository)

	healthController := controllers.NewHealthController(versionService)
	signOnController := controllers.NewSignOnController(userService)
	logInController := controllers.NewLogInController(authService)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.GET("/", healthController.Handle)
	api.POST("/users", signOnController.Handle)
	api.POST("/login", logInController.Handle)
	e.Logger.Fatal(e.Start(host + ":" + port))
}
