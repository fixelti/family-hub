package user

import (
	"context"
	"errors"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (user userUsecase) SignIn(ctx context.Context, email, password string) (models.Tokens, error) {
	foundUser, err := user.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		user.logger.Error("failed to get user by email", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	if foundUser.ID == 0 {
		return models.Tokens{}, customError.ErrInvalidCredentials
	}
 
	if err := user.crypto.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return models.Tokens{}, customError.ErrInvalidCredentials
		}
		user.logger.Error("failed to compare hash and password", zap.Error(err))
		return models.Tokens{}, customError.ErrInternal
	}

	// генерация access токена
	tokens, err := user.jwtLib.GenerateTokens(foundUser.ID)
	if err != nil {
		user.logger.Error("failed to generate jwt tokens", zap.Error(err))
		return tokens, customError.ErrInternal
	}

	return tokens, nil
}
