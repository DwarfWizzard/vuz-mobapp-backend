package rest

import (
	"strings"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/validator"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/repository"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login - POST /auth
func (h *UserHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	mimeType := c.Request().Header.Get(echo.HeaderContentType)
	if !strings.HasPrefix(mimeType, echo.MIMEApplicationJSON) {
		return response.ErrMimeTypeError
	}

	var rq LoginRequest
	if err := c.Bind(&rq); err != nil {
		return response.ErrFormatError
	}

	if !validator.IsEmail(rq.Email) {
		return response.ErrInvalidInput
	}

	user, err := h.uc.AuthorizeUser(ctx, rq.Email, rq.Password)
	if err != nil {
		h.logger.Error("Authroize user by email and password error", zap.Error(err))
		if repository.ErrorIsNoRows(err) {
			return response.ErrUserNotFound
		}

		return response.ErrInternalServerError
	}

	tokenPair, err := h.authService.GenerateUserTokenPair(ctx, user)
	if err != nil {
		h.logger.Error("Generate user token pair error", zap.Error(err))
	}

	return response.Success(c, tokenPair)
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// Refresh - POST /auth/refresh
func (h *UserHandler) Refresh(c echo.Context) error {
	ctx := c.Request().Context()

	mimeType := c.Request().Header.Get(echo.HeaderContentType)
	if !strings.HasPrefix(mimeType, echo.MIMEApplicationJSON) {
		return response.ErrMimeTypeError
	}

	var rq RefreshRequest
	if err := c.Bind(&rq); err != nil {
		return response.ErrFormatError
	}

	tokenPair, err := h.authService.RefreshToken(ctx, rq.RefreshToken)
	if err != nil {
		h.logger.Error("Generate user token pair error", zap.Error(err))
	}

	return response.Success(c, tokenPair)
}
