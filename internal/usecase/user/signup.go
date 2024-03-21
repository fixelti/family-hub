package user

import (
	"context"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"go.uber.org/zap"
)

func (user userUsecase) SignUp(ctx context.Context, email, password string) (models.UserDTO, error) {
	foundUser, err := user.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		user.logger.Error("failed to get user by email", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}
	if foundUser.ID != 0 {
		return models.UserDTO{}, customError.ErrUserExists
	}

	passwordHash, err := user.crypto.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		user.logger.Error("failed to generate from password", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}

	createdUser, err := user.userRepository.Create(ctx, email, string(passwordHash))
	if err != nil {
		user.logger.Error("failed to create user", zap.Error(err))
		return models.UserDTO{}, customError.ErrInternal
	}

	return createdUser.ToUserDTO(), nil
}
