package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Wrapper struct {
	Status string    `json:"status"`
	Data   any       `json:"data,omitempty"`
	Error  *ApiError `json:"error,omitempty"`
}

func (r *Wrapper) Send(c echo.Context) error {
	if r == nil {
		return nil
	}

	var httpCode int
	if r.Status == "error" {
		httpCode = r.Error.Code
	} else {
		httpCode = http.StatusOK
	}

	if c.Request().Method == http.MethodHead {
		return c.NoContent(httpCode)
	} else {
		return c.JSON(httpCode, r)
	}
}

func Success(c echo.Context, data any) error {
	resp := &Wrapper{Status: "success", Data: data}
	return resp.Send(c)
}

func Error(c echo.Context, err *ApiError) error {
	resp := &Wrapper{Status: "error", Error: err}
	return resp.Send(c)
}
