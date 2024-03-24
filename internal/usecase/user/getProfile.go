package user

import (
	"context"

	customError "github.com/fixelti/family-hub/internal/common/errors"
	"github.com/fixelti/family-hub/internal/common/models"
	"go.uber.org/zap"
)

func (user userUsecase) GetProfile(ctx context.Context, userID uint) (models.UserProfile, error) {
	foundUser, err := user.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		user.logger.Error("failed to get user by id", zap.Error(err))
		return models.UserProfile{}, customError.ErrInternal
	}

	if foundUser.ID == 0 {
		return models.UserProfile{}, customError.ErrInvalidCredentials
	}

	services, err := user.diskSASRepository.GetUserServices(ctx, foundUser.ID)
	if err != nil {
		user.logger.Error("failed to get user services: ", zap.Error(err))
		return models.UserProfile{}, customError.ErrInternal
	}

	return models.UserProfile{
		Email:                      foundUser.Email,
		DiskSpaceAllocationService: services,
	}, nil
}
