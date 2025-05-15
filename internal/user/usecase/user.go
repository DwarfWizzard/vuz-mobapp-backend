package usecase

import (
	"context"

	userdto "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usercore"
	"go.uber.org/zap"
)

type UserRepository interface {
	GetUserById(userId uint32) (*usercore.User, error)
	GetUserByEmail(email, pwdHash string) (*usercore.User, error)
	UpdateUser(user *usercore.User) error
}

type GroupProvider interface {
	GetGroupByUserId(ctx context.Context, userId uint32) ([]userdto.UserGroupInfo, error)
}

type UserUseCase struct {
	repo UserRepository

	groupProvider GroupProvider

	logger *zap.Logger
}

func NewUserUC(repo UserRepository, groupProvider GroupProvider, logger *zap.Logger) *UserUseCase {
	return &UserUseCase{
		repo:          repo,
		groupProvider: groupProvider,
		logger:        logger,
	}
}
