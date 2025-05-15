package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInternalServerError  *echo.HTTPError = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrInvalidAuthMethod    *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "invalid authorization method")
	ErrInvalidSigningMethod *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "invalid token signing method")
	ErrMimeTypeError        *echo.HTTPError = echo.NewHTTPError(http.StatusUnsupportedMediaType, "expected content type is application/json.")
	ErrInvalidToken         *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "invalid token")
	ErrInvalidInput                         = echo.NewHTTPError(http.StatusBadRequest, "invalid input")
	ErrFormatError          *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "format error. request is unparsable")
	ErrNotAuthorized        *echo.HTTPError = echo.NewHTTPError(http.StatusForbidden, "not have permission to access the requested data")
	ErrEmailAlreadyInUsed   *echo.HTTPError = echo.NewHTTPError(http.StatusBadRequest, "email is already used")
	ErrUserNotFound         *echo.HTTPError = echo.NewHTTPError(http.StatusNotFound, "user not found")
)
