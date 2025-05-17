package rest

import (
	"errors"
	"log"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/auth"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type GroupScheduleRequest struct {
	GroupId uint32 `param:"id"`
	Date    string `query:"date"`
}

// GroupSchedule - GET /v1/edu/groups/:id/schedule?date=
func (h *EducationHandler) GroupSchedule(c echo.Context) error {
	ctx := c.Request().Context()

	var rq GroupScheduleRequest
	if err := c.Bind(&rq); err != nil {
		return response.ErrFormatError
	}

	apikey, ok := c.Get("apikey").(*auth.Apikey)
	if !ok {
		return response.ErrNotAuthorized
	}

	date, err := time.Parse(time.DateOnly, rq.Date)
	if err != nil {
		h.logger.Error("Parse time error", zap.Error(err))
		return response.ErrInvalidInput
	}

	now := time.Now().Truncate(24 * time.Hour)
	days := now.Sub(date.Truncate(24*time.Hour)).Hours() / 24

	if days > 7 || days < -7 {
		return response.ErrInvalidInput
	}

	log.Println(apikey.UserId)

	schedule, err := h.uc.ListEduGroupSchedule(ctx, apikey.UserId, rq.GroupId, date)
	if err != nil {
		h.logger.Error("List edu group schedule error", zap.Error(err))
		if errors.Is(err, response.ErrNotAuthorized) {
			return response.ErrNotAuthorized
		}

		return response.ErrInternalServerError
	}

	return response.Success(c, schedule)
}
