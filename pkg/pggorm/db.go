package pggorm

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	db *gorm.DB
}

func NewDB(connUrl string, logger *zap.Logger) (*Db, error) {
	db, err := gorm.Open(postgres.Open(connUrl))
	if err != nil {
		return nil, err
	}

	return &Db{db: db}, nil
}

func (db *Db) DB() *gorm.DB {
	return db.db
}

func (db *Db) Close() error {
	sqlDb, err := db.db.DB()
	if err != nil {
		return err
	} else {
		return sqlDb.Close()
	}
}
