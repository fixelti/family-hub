package user

import (
	"context"

	"github.com/fixelti/family-hub/internal/common/models"
	"github.com/fixelti/family-hub/internal/config"
	"github.com/fixelti/family-hub/internal/repository/postgres/diskSpaceAllocationService"
	"github.com/fixelti/family-hub/internal/repository/postgres/user"
	"github.com/fixelti/family-hub/lib/crypto"
	"github.com/fixelti/family-hub/lib/jwt"

	"go.uber.org/zap"
)

type Usecase interface {
	SignUp(ctx context.Context, email, password string) (models.UserDTO, error)
	SignIn(ctx context.Context, email, password string) (models.Tokens, error)
	GetProfile(ctx context.Context, userID uint) (models.UserProfile, error)
}

type userUsecase struct {
	userRepository    user.UserRepository
	diskSASRepository diskSpaceAllocationService.DiskSpaceAllocationServiceRepository
	config            config.Config
	logger            *zap.Logger
	crypto            crypto.Crypto
	jwtLib            jwt.JWT
}

func New(
	userRepository user.UserRepository,
	diskSASRepository diskSpaceAllocationService.DiskSpaceAllocationServiceRepository,
	config config.Config,
	logger *zap.Logger,
	crypto crypto.Crypto,
	jwtLib jwt.JWT) Usecase {
	return userUsecase{
		userRepository:    userRepository,
		diskSASRepository: diskSASRepository,
		config:            config,
		logger:            logger,
		crypto:            crypto,
		jwtLib:            jwtLib,
	}
}
