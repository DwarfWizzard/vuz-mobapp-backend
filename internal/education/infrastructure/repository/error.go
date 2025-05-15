package repository

import (
	"strings"

	"gorm.io/gorm"
)

func ErrorIsNoRows(err error) bool {
	return strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error())
}
