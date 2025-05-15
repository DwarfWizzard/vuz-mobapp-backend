package usecase

import (
	"context"

	"github.com/DwarfWizzard/vuz-mobapp-backend/internal/user/usercore"
	"go.uber.org/zap"
)

func (uc *UserUseCase) AuthorizeUser(ctx context.Context, email, password string) (*usercore.User, error) {
	pwdHash := usercore.PasswordHash(password)

	user, err := uc.repo.GetUserByEmail(email, pwdHash)
	if err != nil {
		uc.logger.Error("Get user by email error", zap.Error(err))
		return nil, err
	}

	return user, nil
}
