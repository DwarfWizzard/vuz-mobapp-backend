package repository

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"
)

type Repo struct {
	dbClient *pggorm.Db
}

// NewRepo
func NewRepo(dbClient *pggorm.Db) *Repo {
	return &Repo{dbClient: dbClient}
}
