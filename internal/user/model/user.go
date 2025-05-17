package model

import (
	"crypto/sha512"
	"encoding/hex"
)

const (
	USER_PASSWORD_SALT = "bTl90E8dlrHdZqRnjf3jZLa7AE8u56Iz"
)

type User struct {
	ID           uint32 `gorm:"primaryKey;type:int;autoIncrement"`
	RoleId       uint8  `gorm:"not null"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string `gorm:"uniqueIndex"`
	DeviceToken  string `gorm:"not null"`
	AvatarURL    string `gorm:"not null"`

	Role *Role `gorm:"foreignKey:RoleId"`
}

type Role struct {
	ID   uint32 `gorm:"primaryKey;type:smallint;autoIncrement"`
	Name string `gorm:"uniqueIndex;not null"`
}

func PasswordHash(pwd string) string {
	hash := sha512.Sum512([]byte(pwd + USER_PASSWORD_SALT))
	return hex.EncodeToString(hash[:])
}

func (u *User) IsStudent() bool {
	return u.Role.Name == "student"
}
