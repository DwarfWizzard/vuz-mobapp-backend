package usecase

import (
	"context"

	userdto "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
)

func (us EducationUseCase) ListGroupsByUserId(ctx context.Context, userId uint32) ([]userdto.UserGroupInfo, error) {
	groups, err := us.repo.ListGroupsByUserId(userId)
	if err != nil {
		return nil, err
	}

	var groupsInfo []userdto.UserGroupInfo
	for _, group := range groups {
		groupsInfo = append(groupsInfo, userdto.UserGroupInfo{
			GroupId:     group.ID,
			GroupNumber: group.Number,
			Faculty:     group.Faculty.Name,
		})
	}

	return groupsInfo, nil
}
