package usercore

import (
	"crypto/sha512"
	"encoding/hex"
)

const (
	USER_PASSWORD_SALT = "bTl90E8dlrHdZqRnjf3jZLa7AE8u56Iz"
)

type User struct {
	ID           uint32 `gorm:"primaryKey;type:serial"`
	Role         *Role  `gorm:"foreignKey:RoleId"`
	RoleId       uint8
	Name         string
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string `gorm:"uniqueIndex"`
	DeviceToken  string
	AvatarURL    string
}

type Role struct {
	Id   uint8  `gorm:"column:id;primaryKey;type:smallserial"`
	Name string `gorm:"uniqueIndex"`
}

func PasswordHash(pwd string) string {
	hash := sha512.Sum512([]byte(pwd + USER_PASSWORD_SALT))
	return hex.EncodeToString(hash[:])
}

func (u *User) IsStudent() bool {
	return u.Role.Name == "student"
}
