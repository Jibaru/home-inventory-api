package middlewares

import (
	"github.com/jibaru/home-inventory-api/m/logger"
	"github.com/labstack/echo/v4"
)

type LoggerMiddleware struct {
}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (m *LoggerMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		logger.LogRequest(ctx.Request().Method, ctx.Request().URL.Path)
		return next(ctx)
	}
}
