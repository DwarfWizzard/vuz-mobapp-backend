package migration

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"
)

func Migrate(dbClient *pggorm.Db) error {
	err := dbClient.DB().AutoMigrate(
		&model.User{},
		&model.Role{},
	)

	if err != nil {
		return err
	}

	return nil
}
