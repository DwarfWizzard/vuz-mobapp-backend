package rest

import (
	"context"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/infrastructure/auth"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usecase"
	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/model"

	"go.uber.org/zap"
)

type UserUseCase interface {
	GetUserInfo(ctx context.Context, userId uint32) (*usecase.UserInfo, error)
	AuthorizeUser(ctx context.Context, email, password string) (*model.User, error)
}

type GroupUseCase interface {
	GetGroupByUserId(ctx context.Context, userId uint32)
}

type UserHandler struct {
	uc          UserUseCase
	authService auth.TokenProvider
	logger      *zap.Logger
}

func NewUserHandler(uc UserUseCase, authService auth.TokenProvider, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		uc:          uc,
		authService: authService,
		logger:      logger,
	}
}
