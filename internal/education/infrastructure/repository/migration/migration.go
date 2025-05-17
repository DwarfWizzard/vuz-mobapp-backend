package migration

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"
)

func Migrate(dbClient *pggorm.Db) error {
	err := dbClient.DB().AutoMigrate(
		&model.Direction{},
		&model.Faculty{},
		&model.Grade{},
		&model.Department{},
		&model.EduGroup{},
		&model.UserGroup{},
		&model.Subject{},
		&model.ScheduleTemplate{},
		&model.ScheduleOverride{},
		&model.ScheduleItemType{},
		&model.Location{},
		&model.Teacher{},
		&model.Degree{},
		&model.EventActivity{},
		&model.EventComment{},
		&model.EventSpeaker{},
		&model.SpeakerEvent{},
		&model.Event{},
	)

	if err != nil {
		return err
	}

	return nil
}
