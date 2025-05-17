package repository

import "github.com/DwarfWizzard/vuz-mobapp-backend/internal/education/model"

// GetGroup
func (r *Repo) GetGroup(groupId uint32) (*model.EduGroup, error) {
	group := &model.EduGroup{
		ID: groupId,
	}

	err := r.dbClient.DB().Debug().
		Preload("Faculty").
		Preload("Direction").
		Preload("Grade").
		Preload("Department").
		First(&group).Error

	if err != nil {
		return nil, err
	}

	return group, nil
}

// GetUserGroup
func (r *Repo) GetUserGroup(userId, groupId uint32) (*model.EduGroup, error) {
	group := &model.EduGroup{
		ID: groupId,
	}

	err := r.dbClient.DB().Debug().
		Preload("Faculty").
		Preload("Direction").
		Preload("Grade").
		Preload("Department").
		Joins("JOIN user_groups ug ON edu_groups.id = ug.edu_groups_id").
		Where("ug.users_id = ?", userId).
		First(&group).Error

	return group, err
}

// ListGroupsByUserId
func (r *Repo) ListGroupsByUserId(userId uint32) ([]model.EduGroup, error) {
	var groups []model.EduGroup
	err := r.dbClient.DB().Debug().
		Preload("Faculty").
		Preload("Direction").
		Preload("Grade").
		Preload("Department").
		Joins("JOIN user_groups ug ON edu_groups.id = ug.edu_groups_id").
		Where("ug.users_id = ?", userId).
		Find(&groups).Error

	if err != nil {
		return nil, err
	}

	return groups, nil
}
