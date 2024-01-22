package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/auth/jwt"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/http/middlewares"
	repositories "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/gorm"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/services/aws"
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
	awsAccessKeyID string,
	awsSecretAccessKey string,
	awsRegion string,
	s3BucketName string,
	db *gorm.DB,
) {
	tokenGenerator := jwt.NewJwtGenerator(jwtSecret, jwtDuration)
	fileManager := aws.NewFileManager(awsAccessKeyID, awsSecretAccessKey, awsRegion, s3BucketName)

	assetRepository := repositories.NewAssetRepository(db)
	versionRepository := repositories.NewVersionRepository(db)
	userRepository := repositories.NewUserRepository(db)
	roomRepository := repositories.NewRoomRepository(db)

	assetService := services.NewAssetService(fileManager, assetRepository)
	authService := services.NewAuthService(userRepository, tokenGenerator)
	userService := services.NewUserService(userRepository)
	versionService := services.NewVersionService(versionRepository)
	roomService := services.NewRoomService(roomRepository)

	healthController := controllers.NewHealthController(versionService)
	signOnController := controllers.NewSignOnController(userService)
	logInController := controllers.NewLogInController(authService)
	createRoomController := controllers.NewCreateRoomController(roomService)
	createAssetController := controllers.NewCreateAssetController(assetService)

	needsAuthMiddleware := middlewares.NewNeedsAuthMiddleware(authService)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/v1")
	api.POST("/login", logInController.Handle)
	api.POST("/users", signOnController.Handle)

	authApi := api.Group("", needsAuthMiddleware.Process)
	authApi.GET("/", healthController.Handle)
	authApi.POST("/rooms", createRoomController.Handle)
	authApi.POST("/assets", createAssetController.Handle)

	e.Logger.Fatal(e.Start(host + ":" + port))
}
