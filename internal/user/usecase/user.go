package usecase

import (
	"context"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"

	"go.uber.org/zap"
)

func NewUserUC(repo UserRepository, groupProvider GroupProvider, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{
		repo:          repo,
		groupProvider: groupProvider,
		logger:        logger,
	}
}

// AuthorizeUser return user by email and password hash
func (uc *UserUseCase) AuthorizeUser(ctx context.Context, email, password string) (*model.User, error) {
	pwdHash := model.PasswordHash(password)

	user, err := uc.repo.GetUserByEmail(email, pwdHash)
	if err != nil {
		uc.logger.Error("Get user by email error", zap.Error(err))
		return nil, err
	}

	return user, nil
}

type UserInfo struct {
	User   *model.User
	Groups []dto.UserGroupInfo
}

// GetUserInfo collects user and his groups into UserInfo
func (uc *UserUseCase) GetUserInfo(ctx context.Context, userId uint32) (*UserInfo, error) {
	user, err := uc.repo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	userGroups, err := uc.groupProvider.ListGroupsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		User:   user,
		Groups: userGroups,
	}, nil
}
