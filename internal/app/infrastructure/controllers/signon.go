package controllers

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SignOnController struct {
	userService *services.UserService
}

type SignOnRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignOnController(userService *services.UserService) *SignOnController {
	return &SignOnController{
		userService,
	}
}

func (c *SignOnController) Handle(ctx echo.Context) error {
	request := SignOnRequest{}

	err := (&echo.DefaultBinder{}).BindBody(ctx, &request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	_, err = c.userService.CreateUser(request.Email, request.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, responses.NewMessageResponse("sign on successfully"))
}
