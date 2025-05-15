package repository

import (
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usercore"
	"github.com/DwarfWizzard/vuz-mobapp-backend/pkg/pggorm"
)

type Repo struct {
	dbClient *pggorm.Db
}

// NewRepo
func NewRepo(dbClient *pggorm.Db) *Repo {
	return &Repo{dbClient: dbClient}
}

// GetUserById
func (r *Repo) GetUserById(userId uint32) (*usercore.User, error) {
	var user *usercore.User
	err := r.dbClient.DB().Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserById
func (r *Repo) GetUserByEmail(email, pwdHash string) (*usercore.User, error) {
	var user *usercore.User
	err := r.dbClient.DB().Where("email = ?", email).Where("password_hash = ?", pwdHash).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserDeviceToken
func (r *Repo) UpdateUser(user *usercore.User) error {
	err := r.dbClient.DB().Select(
		"name",
		"password_hash",
		"device_token",
		"avatar_url",
	).Updates(user).Error

	return err
}
