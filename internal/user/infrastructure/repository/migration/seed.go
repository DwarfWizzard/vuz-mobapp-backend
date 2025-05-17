package migration

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(dbClient *pggorm.Db) error {
	db := dbClient.DB().Clauses(clause.OnConflict{DoNothing: true})

	db.Transaction(func(tx *gorm.DB) error {
		roles := []model.Role{
			{ID: 1, Name: "student"},
		}

		if err := tx.Create(roles).Error; err != nil {
			return err
		}

		users := []model.User{
			{
				ID:           1,
				Name:         "Иван Иванов Борисович",
				Email:        "testmail@host.com",
				PasswordHash: "373c6f9d7ca05bac4f761cd7a6bd7ca9d50a07bd162bee1fec261a4bcb4c0be8e99da542d8cdb948baa074564c13991bd5f31e4cdb8f313f72285b371b9b8b23",
				RoleId:       1,
			},
		}

		if err := tx.Create(users).Error; err != nil {
			return err
		}

		return nil
	})

	return nil
}
