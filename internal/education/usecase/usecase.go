package usecase

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/educore"
	"go.uber.org/zap"
)

type EducationRepository interface {
	GetGroupsByUserId(userId uint32) ([]educore.Group, error)
}

type EducationUseCase struct {
	repo   EducationRepository
	logger *zap.Logger
}

func NewEducationUC(repo EducationRepository, logger *zap.Logger) *EducationUseCase {
	return &EducationUseCase{
		repo:   repo,
		logger: logger,
	}
}
