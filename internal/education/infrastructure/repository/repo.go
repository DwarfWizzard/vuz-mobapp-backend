package repository

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/educore"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"
)

type Repo struct {
	dbClient *pggorm.Db
}

// NewRepo
func NewRepo(dbClient *pggorm.Db) *Repo {
	return &Repo{dbClient: dbClient}
}

func (r *Repo) GetGroupsByUserId(userId uint32) ([]educore.Group, error) {
	var groups []educore.Group
	err := r.dbClient.DB().
		Table("group g").
		Joins("user__group ug ON g.id = ug.group_id").
		Where("ug.user_id = ?", userId).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
