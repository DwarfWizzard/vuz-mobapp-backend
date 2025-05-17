package repository

import "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"

// GetUserById
func (r *Repo) GetUserById(userId uint32) (*model.User, error) {
	var user *model.User
	err := r.dbClient.DB().Preload("Role").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserById
func (r *Repo) GetUserByEmail(email, pwdHash string) (*model.User, error) {
	var user *model.User
	err := r.dbClient.DB().Preload("Role").Where("email = ?", email).Where("password_hash = ?", pwdHash).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUserDeviceToken
func (r *Repo) UpdateUser(user *model.User) error {
	err := r.dbClient.DB().Select(
		"name",
		"password_hash",
		"device_token",
		"avatar_url",
	).Updates(user).Error

	return err
}
