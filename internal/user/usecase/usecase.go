package usecase

import (
	"context"

	userdto "github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/dto"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"

	"go.uber.org/zap"
)

type UserRepository interface {
	GetUserById(userId uint32) (*model.User, error)
	GetUserByEmail(email, pwdHash string) (*model.User, error)
	UpdateUser(user *model.User) error
}

type GroupProvider interface {
	ListGroupsByUserId(ctx context.Context, userId uint32) ([]userdto.UserGroupInfo, error)
}

type UserUseCase struct {
	repo UserRepository

	groupProvider GroupProvider

	logger *zap.Logger
}
