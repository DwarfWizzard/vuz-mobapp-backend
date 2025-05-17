package rest

import (
	"errors"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"go.uber.org/zap"

	"github.com/labstack/echo/v4"
)

// ActiveEvents GET - /v1/edu/events
func (h *EducationHandler) ActiveEvents(c echo.Context) error {
	ctx := c.Request().Context()

	events, err := h.uc.ActiveEvents(ctx)
	if err != nil {
		h.logger.Error("ActiveEvents error", zap.Error(err))
		return response.ErrInternalServerError
	}

	return response.Success(c, events)
}

type ActiveEventsRequest struct {
	EventId uint32 `param:"id"`
}

// ActiveEvents GET - /v1/events/:id
func (h *EducationHandler) EventInfo(c echo.Context) error {
	ctx := c.Request().Context()

	var rq ActiveEventsRequest
	if err := c.Bind(&rq); err != nil {
		return response.ErrFormatError
	}

	event, err := h.uc.GetEventInfo(ctx, rq.EventId)
	if err != nil {
		h.logger.Error("Get event info error", zap.Error(err))
		if errors.Is(err, response.ErrInvalidInput) {
			return response.ErrInvalidInput
		}

		return response.ErrInternalServerError
	}

	return response.Success(c, event)
}
