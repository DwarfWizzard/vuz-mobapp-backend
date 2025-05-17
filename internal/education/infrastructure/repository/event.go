package repository

import (
	"context"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"
)

func (r *Repo) GetEvent(ctx context.Context, eventId uint32) (*model.Event, error) {
	event := &model.Event{
		ID: eventId,
	}

	err := r.dbClient.DB().
		Preload("Faculty").
		Preload("Location").
		Preload("Speakers").
		Preload("Activities").
		Preload("Comments").
		First(&event).Error

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *Repo) ListEvent(ctx context.Context, date string) ([]model.Event, error) {
	var events []model.Event

	err := r.dbClient.DB().
		Select("id", "faculty_id", "title", "date", "time_start", "time_end").
		Preload("Faculty").
		Where("date::date >= ?::date", date).
		Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}
