package rest

import (
	"strings"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/auth"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/logger"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (h *UserHandler) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := logger.Logger()

			ctx := c.Request().Context()

			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer ") {
				return response.ErrInvalidAuthMethod
			}

			tokenValue := authHeader[7:]
			apikey, err := h.authService.GetApikeyByToken(ctx, tokenValue)
			if err != nil {
				log.Error("Get api key error", zap.Error(err))

				if auth.ErrorIsInvalidSigningMethod(err) {
					c.Error(response.ErrInvalidSigningMethod)
				} else if auth.ErrorIsInvalidToken(err) {
					c.Error(response.ErrInvalidToken)
				} else if auth.ErrorIsTokenExpired(err) {
					c.Error(response.ErrTokenExpired)
				} else {
					c.Error(response.ErrInternalServerError)
				}

				return err
			}

			c.Set("apikey", apikey)

			return next(c)
		}
	}
}
