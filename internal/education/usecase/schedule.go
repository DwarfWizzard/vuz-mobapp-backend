package usecase

import (
	"context"
	"time"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/common/response"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/infrastructure/repository"

	"go.uber.org/zap"
)

func (uc *EducationUseCase) ListEduGroupSchedule(ctx context.Context, userId, groupId uint32, date time.Time) ([]dto.Schedule, error) {
	group, err := uc.repo.GetUserGroup(userId, groupId)
	if err != nil {
		uc.logger.Error("Get group error", zap.Error(err))
		if repository.ErrorIsNoRows(err) {
			return nil, response.ErrNotAuthorized
		}

		return nil, err
	}

	isEven := group.IsEvenWeek(date)

	schedule, err := uc.repo.ListSchedule(groupId, date, isEven)
	if err != nil {
		uc.logger.Error("Get schedule error", zap.Error(err))
		return nil, err
	}

	for i, data := range schedule {
		currentWeekday := date.Weekday()

		daysDiff := int(data.WeekDay - currentWeekday)
		data.Date = date.AddDate(0, 0, daysDiff).Format(time.DateOnly)

		schedule[i] = data
	}

	return schedule, nil
}
