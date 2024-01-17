package middlewares

import (
	"github.com/jibaru/home-inventory-api/m/internal/app/application/services"
	"github.com/jibaru/home-inventory-api/m/internal/app/infrastructure/responses"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NeedsAuthMiddleware struct {
	authService *services.AuthService
}

func NewNeedsAuthMiddleware(
	authService *services.AuthService,
) *NeedsAuthMiddleware {
	return &NeedsAuthMiddleware{
		authService,
	}
}

func (m *NeedsAuthMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(
				http.StatusUnauthorized,
				responses.NewMessageResponse("unauthorized"),
			)
		}

		data, err := m.authService.ParseAuthentication(token)
		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				responses.NewMessageResponse(err.Error()),
			)
		}
		c.Set("auth_id", data.ID)

		return next(c)
	}
}
