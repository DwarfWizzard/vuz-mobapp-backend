package rest

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"

	"go.uber.org/zap"
)

type EducationUseCase interface {
	ListEduGroupSchedule(ctx context.Context, userId, groupId uint32, date time.Time) ([]dto.Schedule, error)

	GetEventInfo(ctx context.Context, eventId uint32) (*model.Event, error)
	ActiveEvents(ctx context.Context) ([]model.Event, error)
}

type EducationHandler struct {
	uc     EducationUseCase
	logger *zap.Logger
}

func NewEducationHandler(uc EducationUseCase, logger *zap.Logger) *EducationHandler {
	return &EducationHandler{
		uc:     uc,
		logger: logger,
	}
}
