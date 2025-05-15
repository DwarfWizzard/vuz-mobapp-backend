package response

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func HttpErrorHandler(err error, c echo.Context) {
	log := logger.Logger()
	var eErr *ApiError

	switch e := err.(type) {
	case *echo.HTTPError:
		eErr = &ApiError{
			Code: e.Code,
		}

		if ee, ok := e.Message.(error); ok {
			eErr.Message = ee.Error()
		} else {
			eErr.Message = e.Message.(string)
		}
	case *ApiError:
		eErr = e
	default:
		eErr = &ApiError{
			Code:    -1,
			Message: e.Error(),
		}
	}

	if c.Response().Committed {
		log.Debug("response is commited")
		return
	}

	respErr := Error(c, eErr)

	if respErr != nil {
		log.Warn("can`t send error response")
	}
}
