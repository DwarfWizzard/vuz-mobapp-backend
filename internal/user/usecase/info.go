package usecase

import (
	"context"

	userdto "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usercore"
)

type UserInfo struct {
	User   *usercore.User
	Groups []userdto.UserGroupInfo
}

func (uc *UserUseCase) GetUserInfo(ctx context.Context, userId uint32) (*UserInfo, error) {
	user, err := uc.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	userGroups, err := uc.groupProvider.GetGroupByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		User:   user,
		Groups: userGroups,
	}, nil
}
