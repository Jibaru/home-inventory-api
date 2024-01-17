package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LogInController struct {
	authService *services.AuthService
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LogInResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func NewLogInController(
	authService *services.AuthService,
) *LogInController {
	return &LogInController{
		authService,
	}
}

func (c *LogInController) Handle(ctx echo.Context) error {
	request := LogInRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	data, err := c.authService.Authenticate(request.Email, request.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, responses.NewDataResponse(&LogInResponse{
		ID:    data.User.ID,
		Email: data.User.Email,
		Token: data.Token,
	}))
}
