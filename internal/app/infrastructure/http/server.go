package http

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/controllers"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/http/middlewares"
	repositories "github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/repositories/gorm"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/services/aws"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/services/jwt"
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/labstack/echo/v4"
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
	tokenGenerator := jwt.NewTokenGenerator(jwtSecret, jwtDuration)
	fileManager := aws.NewFileManager(awsAccessKeyID, awsSecretAccessKey, awsRegion, s3BucketName)

	assetRepository := repositories.NewAssetRepository(db)
	versionRepository := repositories.NewVersionRepository(db)
	userRepository := repositories.NewUserRepository(db)
	roomRepository := repositories.NewRoomRepository(db)
	boxRepository := repositories.NewBoxRepository(db)
	itemRepository := repositories.NewItemRepository(db)
	itemKeywordRepository := repositories.NewItemKeywordRepository(db)

	assetService := services.NewAssetService(fileManager, assetRepository)
	authService := services.NewAuthService(userRepository, tokenGenerator)
	userService := services.NewUserService(userRepository)
	versionService := services.NewVersionService(versionRepository)
	roomService := services.NewRoomService(roomRepository, boxRepository)
	boxService := services.NewBoxService(boxRepository, itemRepository, roomRepository)
	itemService := services.NewItemService(itemRepository, itemKeywordRepository, assetService)

	healthController := controllers.NewHealthController(versionService)
	signOnController := controllers.NewSignOnController(userService)
	logInController := controllers.NewLogInController(authService)
	createRoomController := controllers.NewCreateRoomController(roomService)
	createAssetController := controllers.NewCreateAssetController(assetService)
	createBoxController := controllers.NewCreateBoxController(boxService)
	createItemController := controllers.NewCreateItemController(itemService)
	addItemIntoBoxController := controllers.NewAddItemIntoBoxController(boxService)
	removeItemFromBoxController := controllers.NewRemoveItemFromBoxController(boxService)
	getRoomsController := controllers.NewGetRoomsController(roomService)
	getBoxesController := controllers.NewGetBoxesController(boxService)
	getItemsController := controllers.NewGetItemsController(assetService, itemService)
	transferItemController := controllers.NewTransferItemController(boxService)
	deleteBoxController := controllers.NewDeleteBoxController(boxService)
	deleteRoomController := controllers.NewDeleteRoomController(roomService)
	updateRoomController := controllers.NewUpdateRoomController(roomService)
	updateBoxController := controllers.NewUpdateBoxController(boxService)
	updateItemController := controllers.NewUpdateItemController(itemService)
	changeBoxRoomController := controllers.NewChangeBoxRoomController(boxService)
	getBoxTransactionsController := controllers.NewGetBoxTransactionsController(boxService)

	loggerMiddleware := middlewares.NewLoggerMiddleware()
	needsAuthMiddleware := middlewares.NewNeedsAuthMiddleware(authService)

	e := echo.New()
	e.Use(loggerMiddleware.Process)

	api := e.Group("/api/v1")
	api.POST("/login", logInController.Handle)
	api.POST("/users", signOnController.Handle)

	authApi := api.Group("", needsAuthMiddleware.Process)
	authApi.GET("/", healthController.Handle)
	authApi.POST("/rooms", createRoomController.Handle)
	authApi.POST("/rooms/:roomID/boxes", createBoxController.Handle)
	authApi.POST("/assets", createAssetController.Handle)
	authApi.POST("/items", createItemController.Handle)
	authApi.POST("/boxes/:boxID/items", addItemIntoBoxController.Handle)
	authApi.DELETE("/boxes/:boxID/items/:itemID", removeItemFromBoxController.Handle)
	authApi.GET("/rooms", getRoomsController.Handle)
	authApi.GET("/boxes", getBoxesController.Handle)
	authApi.GET("/items", getItemsController.Handle)
	authApi.POST("/boxes/:boxID/items/:itemID/transfer", transferItemController.Handle)
	authApi.DELETE("/boxes/:boxID", deleteBoxController.Handle)
	authApi.DELETE("/rooms/:roomID", deleteRoomController.Handle)
	authApi.PATCH("/rooms/:roomID", updateRoomController.Handle)
	authApi.PATCH("/boxes/:boxID", updateBoxController.Handle)
	authApi.PATCH("/items/:itemID", updateItemController.Handle)
	authApi.PUT("/boxes/:boxID/room", changeBoxRoomController.Handle)
	authApi.GET("/boxes/:boxID/transactions", getBoxTransactionsController.Handle)

	logger.LogError(e.Start(host + ":" + port))
}
