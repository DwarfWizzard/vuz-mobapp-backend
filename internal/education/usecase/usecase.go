package usecase

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"

	"go.uber.org/zap"
)

type EducationRepository interface {
	GetGroup(groupId uint32) (*model.EduGroup, error)
	GetUserGroup(userId, groupId uint32) (*model.EduGroup, error)
	ListGroupsByUserId(userId uint32) ([]model.EduGroup, error)

	ListSchedule(groupId uint32, date time.Time, isEvenWeek bool) ([]dto.Schedule, error)

	GetEvent(ctx context.Context, eventId uint32) (*model.Event, error)
	ListEvent(ctx context.Context, date string) ([]model.Event, error)
}

type EducationUseCase struct {
	repo   EducationRepository
	loc    *time.Location
	logger *zap.Logger
}

func NewEducationUC(repo EducationRepository, loc *time.Location, logger *zap.Logger) *EducationUseCase {
	return &EducationUseCase{
		repo:   repo,
		loc:    loc,
		logger: logger,
	}
}
