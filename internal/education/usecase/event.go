package usecase

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/infrastructure/repository"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"
	"go.uber.org/zap"
)

func (uc *EducationUseCase) GetEventInfo(ctx context.Context, eventId uint32) (*model.Event, error) {
	event, err := uc.repo.GetEvent(ctx, eventId)
	if err != nil {
		uc.logger.Error("Get event error", zap.Error(err))
		if repository.ErrorIsNoRows(err) {
			return nil, response.ErrInvalidInput
		}

		return nil, err
	}

	return event, nil
}

func (uc *EducationUseCase) ActiveEvents(ctx context.Context) ([]model.Event, error) {
	now := time.Now().In(uc.loc)

	events, err := uc.repo.ListEvent(ctx, now.Format(time.DateOnly))
	if err != nil {
		uc.logger.Error("Get list event from current time error", zap.Error(err))
		return nil, err
	}

	return events, nil
}
